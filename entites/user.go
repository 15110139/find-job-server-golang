package entites

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	gorm.Model
	UserId    uuid.UUID `json:"userId"`
	Password  string    `json:"password"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	IsActive  bool      `json:isActive`
}

type Users []User
