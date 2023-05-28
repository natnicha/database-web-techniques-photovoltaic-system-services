package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Help(context *gin.Context) {
	fmt.Println("This is a helper function")
	context.JSON(http.StatusCreated, gin.H{"user": "root"})
}
