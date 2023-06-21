package controller

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"os/exec"
	"path/filepath"
	"photovoltaic-system-services/product/repositories"
	"strconv"

	"github.com/jordan-wright/email"

	"github.com/gin-gonic/gin"
)

type ResponseUser struct {
	Data ResponseData `json:"data"`
}

type ResponseData struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte `json:"password"`
}

func GenerateReport(context *gin.Context) {
	productId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, _ := context.Get("user-id")
	projectNum := repositories.CheckExistProduct(productId)
	if projectNum == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "No specified product ID"})
		return
	}
	productNum := repositories.CheckExistProjectByProductIdAndUserId(productId, userId.(int))
	if productNum == 0 {
		context.JSON(http.StatusConflict, gin.H{"error": "a project ID doesn't belong to a user ID"})
		return
	}

	err = callReportGenerationBatch(context.Param("id"))
	if err != nil {
		log.Println(err.Error())
		context.JSON(http.StatusConflict, gin.H{"error": "generating a report failed: " + err.Error()})
		return
	}
	// TODO: set report MetaData (filename: solar_model_geolocation, get user's email)
	authorization, _ := context.Get("authorization")
	responseUser, err := getUserInfo(authorization.(string))
	if err != nil {
		log.Println(err.Error())
		context.JSON(http.StatusConflict, gin.H{"error": "get user information failed" + err.Error()})
		return
	}
	err = sendEmail(responseUser)
	if err != nil {
		log.Println(err.Error())
		context.JSON(http.StatusConflict, gin.H{"error": "sending an email failed" + err.Error()})
		return
	}
	context.JSON(http.StatusAccepted, nil)
	return
}

func callReportGenerationBatch(productId string) error {
	pythonPath := os.Getenv("PYTHON_PATH")
	generateReportExec := os.Getenv("PYTHON_GENERATE_REPORT_EXEC")
	cmd := exec.Command(pythonPath, generateReportExec, productId)
	return cmd.Run()
}

func getUserInfo(authorization string) (ResponseUser, error) {
	client := &http.Client{}
	url := "http://localhost:" + os.Getenv("SERVICE_PORT") + "/api/v1/user"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseUser{}, err
	}

	request.Header.Set("Authorization", authorization)
	resp, err := client.Do(request)
	if err != nil {
		return ResponseUser{}, err
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
	attachedFiel := find("./", ".xlsx")
	e.AttachFile(attachedFiel[0])
	e.From = from
	e.To = to
	e.Subject = "Photovoltaic System - Report"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := e.Send(smtpHost+":"+smtpPort, auth)
	if err != nil {
		return err
	}
	deleteFile(attachedFiel[0])
	return nil
}

func find(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

func deleteFile(filename string) {
	e := os.Remove(filename)
	if e != nil {
		log.Fatal(e)
	}
}
