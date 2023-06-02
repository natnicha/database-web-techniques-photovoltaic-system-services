package controller

import (
	"net/http"
	"net/url"
	"photovoltaic-system-services/project/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type listAccountRequest struct {
	limit  int32 `form:"limit" binding:"required,min=1"`
	offset int32 `form:"offset" binding:"required,min=0"`
}

func Get(context *gin.Context) {
	var req listAccountRequest
	if err := context.ShouldBindQuery(&req); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	queryParams := context.Request.URL.Query()
	query, err := getQuery(queryParams)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	projects, err := repositories.GetProject(query)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": projects})
}

func getQuery(queryParam url.Values) (query repositories.ListRequest, err error) {
	if len(queryParam["limit"]) > 0 {
		limit, err := strconv.Atoi(queryParam["limit"][0])
		if err != nil {
			return repositories.ListRequest{}, err
		}
		query.Limit = limit
	}
	if len(queryParam["offset"]) > 0 {
		offset, err := strconv.Atoi(queryParam["offset"][0])
		if err != nil {
			return repositories.ListRequest{}, err
		}
		query.Offset = offset
	}
	if len(queryParam["filter"]) > 0 {
		query.Filter = queryParam["filter"][0]
	}
	if len(queryParam["sort_by"]) > 0 {
		query.SortBy = queryParam["sort_by"][0]
	}
	if len(queryParam["order_by"]) > 0 {
		query.OrderBy = queryParam["order_by"][0]
	}
	return
}
