package searchcontrollers

import (
	"strconv"

	_ "fmt"

	service "github.com/find-job-server-golang/service"
	constant "github.com/find-job-server-golang/util/constant"
	response "github.com/find-job-server-golang/util/response"
	"github.com/gin-gonic/gin"
)

type SearchControllers struct {
}

func (searchControllers *SearchControllers) Search(c *gin.Context) {
	text, _ := c.GetQuery("text")
	page, _ := c.GetQuery("page")
	limit, _ := c.GetQuery("limit")
	pageInt, err1 := strconv.Atoi(page)
	if err1 != nil {
		response.RespondWithError(c, constant.PAGE_MUST_BE_NUMBER, 500)
		return
	}
	limitInt, err2 := strconv.Atoi(limit)
	if err2 != nil {
		response.RespondWithError(c, constant.LIMIT_MUST_BE_NUMBER, 500)
		return
	}
	searchService := service.SearchService{}
	dataSearch := searchService.Search(text, pageInt, limitInt)
	response.RespondSuccess(c, gin.H{
		"data": dataSearch,
	}, 200)
	return
}
