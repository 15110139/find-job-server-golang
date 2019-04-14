package comapnycontrolers

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/find-job-server-golang/service"
	"github.com/gin-gonic/gin"
	entities "github.com/find-job-server-golang/entites"
	constant "github.com/find-job-server-golang/util/constant"
	response "github.com/find-job-server-golang/util/response"

)



type CompanyControllers struct {

}

func (companyControllers *CompanyControllers) CreateCompany(c *gin.Context){

	var companyObj entities.Company
	c.ShouldBindJSON(&companyObj)
	companyService := service.CompanyService{}
	_, isNotFound := companyService.FindCompanyWithName(companyObj.Name)
	if !isNotFound{
		response.RespondWithError(c,constant.COMPANY_ALREADY_EXISTS,500)
		return
	}else{
		newCompany:= companyService.CreateCompany(companyObj)
		response.RespondSuccess(c,gin.H{
			"company":newCompany,
		},200)
	}

}


func (companyControllers *CompanyControllers) UpdateCompany(c *gin.Context){

	var dataCompanyUpdate entities.Company
	c.ShouldBindJSON(&dataCompanyUpdate)
	companyService := service.CompanyService{}
	_, isNotFound := companyService.FindCompanyWithID(dataCompanyUpdate.CompanyId)
	if isNotFound{
		response.RespondWithError(c,constant.COMPANY_NOT_FOUND,500)
		return
	}else{
		result:= companyService.UpdateCompanyWithID(dataCompanyUpdate.CompanyId,dataCompanyUpdate)
		response.RespondSuccess(c,gin.H{
			"company":result,
		},200)
	}

}

// func (companyControllers *CompanyControllers) RemoveCompany(c *gin.Context){

// 	var companyId uuid.UUID
// 	c.ShouldBindJSON(&companyId)
// 	companyService := service.CompanyService{}
// 	_, isNotFound := companyService.FindCompanyWithID(companyId)
// 	if isNotFound{
// 		response.RespondWithError(c,constant.COMPANY_NOT_FOUND,500)
// 		return
// 	}else{
// 		result:= companyService.Re(dataCompanyUpdate.CompanyId,dataCompanyUpdate)
// 		response.RespondSuccess(c,gin.H{
// 			"company":result,
// 		},200)
// 	}

// }


func (companyControllers *CompanyControllers) Companies(c *gin.Context){
	companyService := service.CompanyService{}
	companies:= companyService.Companies()
	response.RespondSuccess(c,gin.H{
		"data":companies,
	},200)
	return
}


func (companyControllers *CompanyControllers) Comapny(c *gin.Context){
	companyService := service.CompanyService{}
	companyId,_:= uuid.FromString(c.Params.ByName("companyId"))
	fmt.Println(companyId)
	company,isNotFound:= companyService.FindCompanyWithID(companyId)
	if(isNotFound){
		response.RespondWithError(c,constant.COMPANY_NOT_FOUND,500)
		return
	}else{
		response.RespondSuccess(c,gin.H{
			"data":company,
		},200)
		return
	}
	
}
