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
	CompanyId    uuid.UUID  `json:"companyId"`
	Name         string   `json:"name"`
	// Technologies []string `gorm: "technologies type:string[]"`
	Technologies string `json:"technologies"`
	CompanyType  string   `json:"companyType"`
	Location     string   `json:"location"`
	// Decs         string   `json:"decs"`
	// OurStore     string   `json:"ourStore"`
	// OurPeople    []People `gorm: "type:People[]"`
}
