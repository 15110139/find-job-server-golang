package main

import (
	router "github.com/find-job-server-golang/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	router := router.Router{}

	router.Auth(v1.Group("/auth"))
	router.Search(v1)
	// apiv1.GET("/v1/search", searchControllerr.Search)
	r.Run(":3000")
}
