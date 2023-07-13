package controller

import (
	"fmt"
	"net/http"
	"net/url"
	"photovoltaic-system-services/project/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(context *gin.Context) {
	queryParams := context.Request.URL.Query()

	userIdStr, _ := context.Get("user-id")
	userId := fmt.Sprint(userIdStr.(int))
	if queryParams.Get("filter") != "" {
		queryParams.Set("filter", queryParams.Get("filter")+"&user_id:"+userId)
	} else {
		queryParams.Add("filter", "user_id:"+userId)
	}

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
