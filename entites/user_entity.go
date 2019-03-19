package entities

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	gorm.Model
	UserId    uuid.UUID
	UserName  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type Users []User
