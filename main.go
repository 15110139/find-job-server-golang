package main

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/find-job-server-golang/controllers"
	// "github.com/find-job-server-golang/middleware"
)

func main() {
	userControllerr := controllers.UserControllers{}
	searchControllerr := controllers.SearchControllers{}
	r := gin.Default()
	r.POST("/register", userControllerr.Register)
	r.POST("/login", userControllerr.Login)
	r.GET("/search", searchControllerr.Search)
	r.Run(":6969") // listen and serve on 0.0.0.0:8080
}
