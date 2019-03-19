package main

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/testgolang/controllers"
)

func main() {
	userControllerr := controllers.UserControllers{}
	r := gin.Default()
	r.POST("/register", userControllerr.Register)
	r.POST("/login", userControllerr.Login)
	r.GET("/parseToken", userControllerr.ParseToken)
	r.Run() // listen and serve on 0.0.0.0:8080
}
