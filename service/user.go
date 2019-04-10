package service

import (
	"fmt"

	"github.com/find-job-server-golang/config"
	entities "github.com/find-job-server-golang/entites"
	uuid "github.com/satori/go.uuid"
)

type UserService struct {
}

func (userService *UserService) CreateUser(user entities.User) entities.User {
	db := config.GetPostgersDB()
	db.AutoMigrate(&entities.User{})
	u1, err := uuid.NewV1()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		panic(err)
	}
	user = entities.User{UserId: u1, Email: user.Email, FirstName: user.FirstName, LastName: user.LastName, Password: user.Password, IsActive: true}
	db.Create(&user)
	return user
}

func (userService *UserService) FindUserWithEmail(email string) (entities.User, bool) {
	db := config.GetPostgersDB()
	db.AutoMigrate(&entities.User{})
	var user entities.User
	isNotFound := db.Where("email = ?", email).First(&user).RecordNotFound()
	return user, isNotFound
}

func (userService *UserService) FindUserWithUserId(userId string) (entities.User, bool) {
	db := config.GetPostgersDB()
	var user entities.User
	isNotFound := db.Where("user_id = ?", userId).First(&user).RecordNotFound()
	return user, isNotFound
}

func (userService *UserService) UpdateProfile(userId string, user entities.User) entities.User {
	db := config.GetPostgersDB()
	var result entities.User
	db.Model(&result).Where("user_id = ?", userId).Update(user)
	return result
}
