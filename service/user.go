package service

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/find-job-server-golang/config"
	entities "github.com/find-job-server-golang/entites"
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
	user = entities.User{UserId: u1, Email: user.Email, FirstName: user.FirstName, LastName: user.LastName, UserName: user.UserName, Password: user.Password}
	db.Create(&user)
	return user
}

func (userService *UserService) FindUserWithUserName(UserName string) entities.User {
	db := config.GetPostgersDB()
	db.AutoMigrate(&entities.User{})
	var user entities.User
	fmt.Printf("username", UserName)
	db.Find(&user, "user_name = ?", UserName)
	return user
}

func (userService *UserService) FindUserWithID(ID string) entities.User {
	db := config.GetPostgersDB()
	// db.AutoMigrate(&entities.User{})
	var user entities.User
	db.Find(&user, "ID = ?", ID)
	return user
}
