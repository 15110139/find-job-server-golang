package controllers

import (
	uuid "github.com/satori/go.uuid"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	entities "github.com/find-job-server-golang/entites"
	service "github.com/find-job-server-golang/service"
	util "github.com/find-job-server-golang/util"
)

type UserControllers struct {
}

type MyCustomClaims struct {
	UserId uuid.UUID `json:"userId"`
	jwt.StandardClaims
}

func (userControllers *UserControllers) Register(c *gin.Context) {
	mySigningKey := []byte("AllYourBase")
	var userObj entities.User
	c.ShouldBindJSON(&userObj)
	userService := service.UserService{}
	user := userService.CreateUser(userObj)
	claims := MyCustomClaims{
		user.UserId,
		jwt.StandardClaims{
			ExpiresAt: 84000,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
		util.RespondWithError(c, "CREATE_TOKEN_FAILURE")
		return
	}else{
		util.RespondSuccess(c, gin.H{
			"token": ss,
		})
		return
	}
}

func (userControllers *UserControllers) Login(c *gin.Context) {
	mySigningKey := []byte("AllYourBase")
	var userObj entities.User
	c.ShouldBindJSON(&userObj)
	var userName = userObj.UserName
	userService := service.UserService{}
	user := userService.FindUserWithUserName(userName)
	if user.Password != userObj.Password {
		util.RespondWithError(c, "INCRRECT_PASSWORD")
		return
	} else {
		claims := MyCustomClaims{
			user.UserId,
			jwt.StandardClaims{
				ExpiresAt: 84000,
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)
		if err != nil {
			util.RespondWithError(c, "CREATE_TOKEN_FAILURE")
			return

		}else{
			util.RespondSuccess(c, gin.H{
				"token": ss,
			})
			return
		}
	}
	// Create the Claims
	

}
