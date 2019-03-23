package entites

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type People struct {
	Name      string
	Position  string
	Introduce string
}

type Company struct {
	gorm.Model
	CompanyId    uuid.UUID
	Name         string   `json:"name"`
	Technologies []string `gorm: "type:string[]"`
	CompanyType  string   `json:"companyType"`
	Location     string   `json:"location"`
	Decs         string   `json:"decs"`
	OurStore     string   `json:"ourStore"`
	OurPeople    []People `gorm: "type:People[]"`
}
