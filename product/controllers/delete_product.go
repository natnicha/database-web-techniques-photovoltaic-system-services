package controller

import (
	"net/http"
	"photovoltaic-system-services/product/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Delete(context *gin.Context) {
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
	err = repositories.DeleteProductById(productId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, nil)
}
