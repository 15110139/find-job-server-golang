package comapnycontrolers

import (
	"strconv"
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
	page,_ :=c.GetQuery("page")
	limit,_:=c.GetQuery("limit")
	pageInt, err1 := strconv.Atoi(page)
	if err1 != nil {
		response.RespondWithError(c, constant.PAGE_MUST_BE_NUMBER, 500)
		return
	}
	limitInt, err2 := strconv.Atoi(limit)
	if err2 != nil {
		response.RespondWithError(c, constant.LIMIT_MUST_BE_NUMBER, 500)
		return
	}
	companies := companyService.Companies(limitInt,pageInt)
	allRecordCompany := companyService.CompaniesCount()
	response.RespondSuccess(c,gin.H{
		"data":companies,
		"totalPage":allRecordCompany/limitInt,
		"currentPage":pageInt,
	},200)
	return
}


func (companyControllers *CompanyControllers) Company(c *gin.Context){
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
