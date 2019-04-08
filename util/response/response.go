package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, err string, code int) {
	c.JSON(http.StatusNotFound, gin.H{
		"error":  err,
		"data":   nil,
		"status": "Error",
		"code":   code | 500,
	})
	c.Abort()
}

func RespondSuccess(c *gin.Context, data gin.H, code int) {
	c.JSON(http.StatusOK, gin.H{
		"error":  nil,
		"data":   data,
		"status": "Ok",
		"code":   code | 200,
	})
	return
}
