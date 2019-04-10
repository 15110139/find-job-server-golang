package main

import (
	router "github.com/find-job-server-golang/router"
	"github.com/gin-gonic/gin"
	// "github.com/find-job-server-golang/middleware"

)

func main() {
	r := gin.Default()
	// r.Use(middleware.RequestIdMiddleware())
	v1 := r.Group("/v1")
	router := router.Router{}

	router.Auth(v1.Group("/auth"))
	router.Company(v1.Group("/company"))

	router.Search(v1)
	
	// apiv1.GET("/v1/search", searchControllerr.Search)
	r.Run(":6969")
}
