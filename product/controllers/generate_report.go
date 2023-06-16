package controller

import (
	"net/http"
	"os/exec"
	"photovoltaic-system-services/product/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	report, err := callReportGenerationBatch(productId)
	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"error": "generating a report failed"})
		return
	}
	err = sendEmail(report)
	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"error": "sending an email failed"})
		return
	}
	context.JSON(http.StatusAccepted, nil)
	return
}

func callReportGenerationBatch(productId int) ([]byte, error) {
	return exec.Command("/usr/local/opt/bin/python3.7", "/users/test.py", "-i", "12.13.14.15", "--cmd", "uptime && date").Output()
}

func sendEmail(report []byte) error {
	return nil
}
