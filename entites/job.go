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

type Job struct {
	gorm.Model
	JobId    uuid.UUID  `json:"jobId"`
	Name         string   `json:"name"`
	Technologies pq.StringArray `gorm: "type:text[]"` 
	Location     string   `json:"location"`
	Decs         pq.StringArray `gorm: "type:text[]"` 
	Require      pq.StringArray `gorm: "type:text[]"` 
	IsActive    bool 	  `json:"isActive"`
}
