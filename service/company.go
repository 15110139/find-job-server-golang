package service

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/find-job-server-golang/config"
	entities "github.com/find-job-server-golang/entites"
)

type CompanyService struct {
}

func (companyService *CompanyService) CreateCompany(company entities.Company) entities.Company {
	db := config.GetPostgersDB()
	db.AutoMigrate(&entities.Company{})
	u1, err := uuid.NewV1()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		panic(err)
	}
	company = entities.Company{CompanyId: u1, Name: company.Name, CompanyType: company.CompanyType, Location: company.Location}
	db.Create(&company)
	return company
}

func (companyService *CompanyService) FindCompanyWithID(ID uuid.UUID) (entities.Company,bool) {
	db := config.GetPostgersDB()
	// db.AutoMigrate(&entities.User{})
	var company entities.Company
	isNotFound := db.Where("company_id = ?", ID).First(&company).RecordNotFound()
	return company, isNotFound
}


func (companyService*CompanyService) UpdateCompanyWithID(ID uuid.UUID, company entities.Company) entities.Company{
	db:= config.GetPostgersDB()
	var result entities.Company
	db.Model(&result).Where("company_id = ?",ID).Updates(company)
	return result
}

func (companyService *CompanyService) FindCompanyWithName(name string) (entities.Company, bool) {
	db := config.GetPostgersDB()
	db.AutoMigrate(&entities.Company{})
	var company entities.Company
	isNotFound := db.Where("name = ?", name).First(&company).RecordNotFound()
	return company, isNotFound
}

// func (companyService *CompanyService) RemoveCompanyWithID(ID uuid.UUID) (entities.Company, bool) {
// 	db := config.GetPostgersDB()
// 	db.AutoMigrate(&entities.Company{})
// 	var company entities.Company
// 	isNotFound := db.Where("company_id = ?", ID).Update("is_active")
// 	return company, isNotFound
// }


func (companyService *CompanyService) Companies() []entities.Company{
	db := config.GetPostgersDB()
	db.AutoMigrate(&entities.Company{})
	var companies []entities.Company
	db.Find(&companies)
	return companies
}
