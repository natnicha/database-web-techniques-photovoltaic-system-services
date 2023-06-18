package controller

import (
	"net/http"
	"os"
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

	report, err := callReportGenerationBatch(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"error": "generating a report failed"})
		return
	}
	// TODO: set report MetaData (filename: solar_model_geolocation, get user's email)
	err = sendEmail(report)
	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"error": "sending an email failed"})
		return
	}
	context.JSON(http.StatusAccepted, nil)
	return
}

func callReportGenerationBatch(productId string) ([]byte, error) {
	pythonPath := os.Getenv("PYTHON_PATH")
	generateReportExec := os.Getenv("PYTHON_GENERATE_REPORT_EXEC")
	return exec.Command(pythonPath, generateReportExec, "-i", "12.13.14.15", "--cmd", productId).Output()
}

func sendEmail(report []byte) error {
	return nil
}
