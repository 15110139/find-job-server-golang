package util

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, err string) {
	c.JSON(http.StatusNotFound, gin.H{
		"error":  err,
		"data":   nil,
		"status": "Error",
	})
}

func RespondSuccess(c *gin.Context, data gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"error":  nil,
		"data":   data,
		"status": "Successfull",
	})
}
