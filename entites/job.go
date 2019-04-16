package entites

import (
	"github.com/lib/pq"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)


type Job struct {
	gorm.Model
	JobId    uuid.UUID  `json:"jobId"`
	Name string  `json:"name"`
	Require  pq.StringArray `json:"require" gorm: "type:text[]"` 
	Desc  pq.StringArray `json:"desc" gorm: "type:text[]"` 
	Technologies  pq.StringArray `json:"technologies" gorm: "type:text[]"` 
	IsActive bool `json:"isActive"`
}
