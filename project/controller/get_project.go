package controller

import (
	"net/http"
	"photovoltaic-system-services/project/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type listAccountRequest struct {
	limit  int32 `form:"limit" binding:"required,min=1"`
	offset int32 `form:"offset" binding:"required,min=0"`
}

func Get(context *gin.Context) {
	// TODO set default value
	var req listAccountRequest
	if err := context.ShouldBindQuery(&req); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	query := context.Request.URL.Query()
	limit, err := strconv.Atoi(query["limit"][0])
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	offset, err := strconv.Atoi(query["offset"][0])
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// TODO extract filter
	// var filter repositories.Filter
	// json.Unmarshal([]byte(query["filter"][0]), &filter)

	query2 := repositories.ListRequest{
		Limit:   limit,
		Offset:  offset,
		SortBy:  query["sort_by"][0],
		OrderBy: query["order_by"][0],
	}
	projects, err := repositories.GetProject(query2)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": projects})
}
