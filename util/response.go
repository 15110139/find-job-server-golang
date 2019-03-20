package util

import "github.com/gin-gonic/gin"

func RespondWithError(c *gin.Context, err string) {
	c.AbortWithStatusJSON(500, gin.H{
		"error":  err,
		"data":   nil,
		"status": "Error",
	})
}

func RespondSuccess(c *gin.Context, data gin.H) {
	c.JSON(200, gin.H{
		"error":  nil,
		"data":   data,
		"status": "Successfull",
	})
}
