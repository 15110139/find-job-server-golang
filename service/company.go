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
	company = entities.Company{CompanyId: u1, Name: company.Name, Technologies: company.Technologies, CompanyType: company.CompanyType, Location: company.Location}
	db.Create(&company)
	return company
}

func (companyService *CompanyService) FindCompanyWithID(ID string) entities.Company {
	db := config.GetPostgersDB()
	// db.AutoMigrate(&entities.User{})
	var company entities.Company
	db.Find(&company, "ID = ?", ID)
	return company
}
