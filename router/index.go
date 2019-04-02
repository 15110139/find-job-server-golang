package router

import (
	controllers "github.com/find-job-server-golang/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	auth   controllers.UserControllers
	search controllers.SearchControllers
}

func (auth *Router) Auth(router *gin.RouterGroup) {
	router.POST("/register", auth.auth.Login)
	router.POST("/login", auth.auth.Login)
}

func (search *Router) Search(router *gin.RouterGroup) {
	router.GET("/search", search.search.Search)
}
