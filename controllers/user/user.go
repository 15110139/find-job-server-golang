package usercontrollers

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	entities "github.com/find-job-server-golang/entites"
	service "github.com/find-job-server-golang/service"
	constant "github.com/find-job-server-golang/util/constant"
	response "github.com/find-job-server-golang/util/response"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
	_, isNotFound := userService.FindUserWithEmail(userObj.Email)
	if isNotFound == false {
		response.RespondWithError(c, constant.EMAIL_ALREADY_EXISTS, 500)
		return
	}
	newUser := userService.CreateUser(userObj)
	fmt.Println(newUser)
	claims := MyCustomClaims{
		newUser.UserId,
		jwt.StandardClaims{
			ExpiresAt: 84000,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
		response.RespondWithError(c, constant.CREATE_TOKEN_FAILURE, 500)
		return
	} else {
		response.RespondSuccess(c, gin.H{
			"token":   ss,
			"profile": newUser,
		}, 200)
		return
	}
}

func (userControllers *UserControllers) Login(c *gin.Context) {
	mySigningKey := []byte("AllYourBase")
	var userObj entities.User
	c.ShouldBindJSON(&userObj)
	var email = userObj.Email
	userService := service.UserService{}
	user, isNotFound := userService.FindUserWithEmail(email)
	if isNotFound {
		response.RespondWithError(c, constant.USER_NOT_FOUND, 500)
		return
	}
	if user.Password != userObj.Password {
		response.RespondWithError(c, constant.INCRRECT_PASSWORD, 500)
		return
	} else {

		claims := MyCustomClaims{
			user.UserId,
			jwt.StandardClaims{
				ExpiresAt: 8400000000000000000,
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)
		if err != nil {
			response.RespondWithError(c, constant.CREATE_TOKEN_FAILURE, 500)
			return

		} else {
			response.RespondSuccess(c, gin.H{
				"token":   ss,
				"profile": user,
			}, 200)
			return
		}
	}

}

func (userControllers *UserControllers) UpdateProfile(c *gin.Context) {
	var dataUserUpdate entities.User
	userIdFormToken,_ := c.Get("userIdFormToken")
	c.ShouldBindJSON(&dataUserUpdate)
	userService := service.UserService{}
	var password = dataUserUpdate.Password
	if len(password) > 0 {
		response.RespondWithError(c, constant.INVALID_PARAMETERS, 500)
		return
	}
	_, isNotFound := userService.FindUserWithUserId(userIdFormToken.(string))
	if isNotFound {
		response.RespondWithError(c, constant.USER_NOT_FOUND, 500)
		return
	}
	result := userService.UpdateProfile(userIdFormToken.(string), dataUserUpdate)
	response.RespondSuccess(c, gin.H{
		"data": result,
	}, 200)
	return

}

func ( UserControllers *UserControllers) ValidateToken (c* gin.Context){
	userService := service.UserService{}
	userIdFormToken,_ := c.Get("userIdFormToken")
	user, isNotFound := userService.FindUserWithUserId(userIdFormToken.(string))
	if isNotFound {
		response.RespondWithError(c, constant.USER_NOT_FOUND, 500)
		return
	}else{
		response.RespondSuccess(c, gin.H{
			"profile": user,
		}, 200)
		return
	}
}
