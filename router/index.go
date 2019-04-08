package router

import (
	searchcontrollers "github.com/find-job-server-golang/controllers/search"
	usercontrollers "github.com/find-job-server-golang/controllers/user"
	"github.com/find-job-server-golang/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	auth   usercontrollers.UserControllers
	search searchcontrollers.SearchControllers
	c   gin.Context

}

func (route *Router) Auth(router *gin.RouterGroup) {
	router.POST("/register", route.auth.Register)
	router.POST("/login", route.auth.Login)
	router.POST("/updateProfile", middleware.TokenAuthMiddleware,route.auth.UpdateProfile)
	router.POST("/validateToken", middleware.TokenAuthMiddleware,route.auth.ValidateToken)

}

func (route *Router) Search(router *gin.RouterGroup) {
	router.GET("/search", route.search.Search)
}
