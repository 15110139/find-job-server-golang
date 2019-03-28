package controllers

import (
	"strconv"

	"github.com/testgolang/util"
	_"fmt"
	"github.com/gin-gonic/gin"
	service "github.com/testgolang/service"
)

type SearchControllers struct {
}

func (searchControllers *SearchControllers) Search(c *gin.Context) {
	text, _ := c.GetQuery("text")
	offset, _ := c.GetQuery("offset")
	limit, _ := c.GetQuery("limit")
	offsetInt, err1 := strconv.Atoi(offset)
	if err1 != nil {
		util.RespondWithError(c, "OFFSET_MUST_BE_NUMBER")
		return
	}
	limitInt, err2 := strconv.Atoi(limit)
	if err2 != nil {
		 util.RespondWithError(c, "LIST_MUST_BE_NUMBER")
		 return
	}
	searchService := service.SearchService{}
	dataSearch := searchService.Search(text, offsetInt, limitInt)
	util.RespondSuccess(c, gin.H{
		"data": dataSearch,
	})
	return
}
