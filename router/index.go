package router

import (
	searchcontrollers "github.com/find-job-server-golang/controllers/search"
	usercontrollers "github.com/find-job-server-golang/controllers/user"
	companycontrollers "github.com/find-job-server-golang/controllers/company"
	"github.com/find-job-server-golang/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	auth   usercontrollers.UserControllers
	search searchcontrollers.SearchControllers
	company companycontrollers.CompanyControllers
}

func (route *Router) Auth(router *gin.RouterGroup) {
	router.POST("/register", route.auth.Register)
	router.POST("/login", route.auth.Login)
	router.PUT("/updateProfile", middleware.TokenAuthMiddleware,route.auth.UpdateProfile)
	router.POST("/validateToken", middleware.TokenAuthMiddleware,route.auth.ValidateToken)

}

func (route *Router) Search(router *gin.RouterGroup) {
	router.GET("/search", route.search.Search)
}

func (route *Router) Company(router *gin.RouterGroup) {
	router.POST("/createCompany", route.company.CreateCompany)
	router.PUT("/updateCompany", route.company.UpdateCompany)
	router.GET("/companies", route.company.Companies)
	// router.GET("/company/:companyId", route.company.Company)
}


