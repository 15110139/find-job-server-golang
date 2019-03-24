package main

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/testgolang/controllers"
	"github.com/testgolang/middleware"
)

func main() {

	userControllerr := controllers.UserControllers{}
	searchControllerr := controllers.SearchControllers{}

	r := gin.Default()
	r.POST("/register", middleware.TokenAuthMiddleware, userControllerr.Register)
	r.POST("/login", userControllerr.Login)
	r.GET("/search",searchControllerr.Search)
	r.Run() // listen and serve on 0.0.0.0:8080
}
