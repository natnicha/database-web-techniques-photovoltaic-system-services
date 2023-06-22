package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"photovoltaic-system-services/project/repositories"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
)

type ResponseUser struct {
	Data ResponseUserData `json:"data"`
}

type ResponseUserData struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte `json:"password"`
}

type ResponseGetProduct struct {
	Data []ResponseGetProductData `json:"data"`
}

type ResponseGetProductData struct {
	Id                int     `json:"id"`
	ProjectId         int     `json:"project_id"`
	SolarPanelModelId int     `json:"solar_panel_model_id"`
	Orientation       float32 `json:"orientation"`
	Inclination       float32 `json:"inclination"`
	Area              float32 `json:"area"`
	Geolocation       string  `json:"geolocation"`
}

func GenerateReport(context *gin.Context) {
	projectId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, _ := context.Get("user-id")
	projectNum := repositories.CheckExistProject(projectId, repositories.Projects{UserId: userId.(int)})
	if projectNum == 0 {
		context.JSON(http.StatusConflict, gin.H{"error": "a project ID doesn't belong to a user ID"})
		return
	}

	project, err := repositories.GetProject(repositories.ListRequest{Filter: "id:" + context.Param("id")})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "get project information failed" + err.Error()})
		return
	}

	var projectObj = (*project)
	if projectObj[0].IsPrinted {
		context.JSON(http.StatusConflict, gin.H{"error": "the project was already printed"})
		return
	}
	var wg sync.WaitGroup

	// immediatly return 202 Accepted response
	go func() {
		wg.Add(1)
		defer wg.Done()
		context.JSON(http.StatusAccepted, nil)
		return
	}()

	// parallel making requests to /product/generate-report
	go func() {
		// check data, if previous 30 days data is missing, call history
		authorization, _ := context.Get("authorization")
		wg.Wait()
		products, err := requestProducts(authorization.(string), context.Param("id"))
		if err != nil {
			log.Println(err.Error())
			context.JSON(http.StatusInternalServerError, gin.H{"error": "get product information failed" + err.Error()})
			return
		}

		for _, product := range products.Data {
			wg.Add(1)
			go func(p ResponseGetProductData) {
				defer wg.Done()

				err = requestProductReportGeneration(authorization.(string), p.Id)
				if err != nil {
					log.Println(err.Error())
					context.JSON(http.StatusInternalServerError, gin.H{"error": "generating a report failed: " + err.Error()})
					return
				}
			}(product)
		}
		wg.Wait()

		responseUser, err := getUserInfo(authorization.(string))
		if err != nil {
			log.Println(err.Error())
			context.JSON(http.StatusInternalServerError, gin.H{"error": "get user information failed" + err.Error()})
			return
		}
		err = sendEmail(responseUser)
		if err != nil {
			log.Println(err.Error())
			context.JSON(http.StatusInternalServerError, gin.H{"error": "sending an email failed" + err.Error()})
			return
		}
		projectObj[0].IsPrinted = true
		repositories.UpdateProject(projectId, projectObj[0])
		return
	}()
}

func requestProducts(authorization string, productId string) (ResponseGetProduct, error) {
	client := &http.Client{}
	url := "http://localhost:" + os.Getenv("SERVICE_PORT") + "/api/v1/product/"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseGetProduct{}, err
	}
	q := request.URL.Query()
	q.Add("filter", "project_id:"+productId)
	request.URL.RawQuery = q.Encode()

	request.Header.Set("Authorization", authorization)
	resp, err := client.Do(request)
	if err != nil {
		return ResponseGetProduct{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return ResponseGetProduct{}, errors.New("Internal request failed")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result ResponseGetProduct
	err = json.Unmarshal(body, &result)
	if err != nil {
		return ResponseGetProduct{}, err
	}
	return result, nil
}

func requestProductReportGeneration(authorization string, productId int) error {
	client := &http.Client{}
	url := "http://localhost:" + os.Getenv("SERVICE_PORT") + "/api/v1/product/generate-report/" + fmt.Sprint(productId)
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", authorization)
	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusAccepted {
		return errors.New("Internal request failed")
	}
	return nil
}

func getUserInfo(authorization string) (ResponseUser, error) {
	client := &http.Client{}
	url := "http://localhost:" + os.Getenv("SERVICE_PORT") + "/api/v1/user/"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseUser{}, err
	}

	request.Header.Set("Authorization", authorization)
	resp, err := client.Do(request)
	if err != nil {
		return ResponseUser{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return ResponseUser{}, errors.New("Internal request failed")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result ResponseUser
	err = json.Unmarshal(body, &result)
	if err != nil {
		return ResponseUser{}, err
	}
	return result, nil
}

func sendEmail(userInfo ResponseUser) error {

	from := "nat.rodtong@gmail.com"
	password := "ebhljkytaepcqawa"
	to := []string{
		userInfo.Data.Email,
	}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	e := email.NewEmail()
	attachedFiel := find("./", ".xlsx", userInfo.Data.Id)
	for _, file := range attachedFiel {
		e.AttachFile(file)
	}
	e.From = from
	e.To = to
	e.Subject = "Photovoltaic System - Report"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := e.Send(smtpHost+":"+smtpPort, auth)
	if err != nil {
		return err
	}

	for _, file := range attachedFiel {
		deleteFile(file)
	}
	return nil
}

func find(root, ext string, userId int) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			file := filepath.Base(d.Name())
			fielUserId := strings.Split(file, "-")[0]
			if fielUserId == fmt.Sprint(userId) {
				a = append(a, s)
			}
		}
		return nil
	})
	return a
}

func deleteFile(filename string) {
	e := os.Remove(filename)
	if e != nil {
		log.Println(e)
	}
}
