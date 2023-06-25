package controller

import (
	"log"
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
		context.JSON(http.StatusConflict, gin.H{"error": "a product ID doesn't belong to a user ID"})
		return
	}

	err = callReportGenerationBatch(context.Param("id"))
	if err != nil {
		log.Println(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"error": "generating a report failed: " + err.Error()})
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
