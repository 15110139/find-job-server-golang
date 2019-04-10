package entites

import (
	"github.com/lib/pq"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// type People struct {
// 	PeopleId uuid.UUID `json:"peopleId"`
// 	Name      string
// 	Position  string
// 	Introduce string
// }

type Company struct {
	gorm.Model
	CompanyId    uuid.UUID  `json:"companyId"`
	Name         string   `json:"name"`
	Technologies pq.StringArray `gorm: "type:text[]"` 
	CompanyType  string   `json:"companyType"`
	Location     string   `json:"location"`
	Decs         string   `json:"decs"`
	OurStory    string   `json:"ourStory"`
	// OurPeople    []People `gorm: "type:People[]"`
	IsActive    bool 	  `json:"isActive"`
}
