package controllers

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	entities "github.com/testgolang/entites"
	userservice "github.com/testgolang/service"
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
	fmt.Println(userObj)
	userService := userservice.UserService{}
	user := userService.CreateUser(userObj)
	claims := MyCustomClaims{
		user.UserId,
		jwt.StandardClaims{
			ExpiresAt: 86400,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  "CREATE_TOKE_FAILURE",
			"data":   nil,
			"status": "Error",
		})
	}
	c.JSON(200, gin.H{
		"error": nil,
		"data": gin.H{
			"token": ss,
		},
		"status": "Successfull",
	})
}

func (userControllers *UserControllers) Login(c *gin.Context) {
	mySigningKey := []byte("AllYourBase")
	var userObj entities.User
	c.ShouldBindJSON(&userObj)
	var userName = userObj.UserName
	userService := userservice.UserService{}
	user := userService.FindUserWithUserName(userName)
	if user.Password != userObj.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "INCORRECT_PASSWORD",
			"data":   nil,
			"status": "Error",
		})
		panic("INCORRECT_PASSWORD")
	}
	// Create the Claims
	claims := MyCustomClaims{
		user.UserId,
		jwt.StandardClaims{
			ExpiresAt: 86400,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  "CREATE_TOKE_FAILURE",
			"data":   nil,
			"status": "Error",
		})
	}
	c.JSON(200, gin.H{
		"error": nil,
		"data": gin.H{
			"token": ss,
		},
		"status": "Successfull",
	})
}

func (userControllers *UserControllers) ParseToken(c *gin.Context) {
	mySigningKey := []byte("AllYourBase")
	tokenString := c.Request.Header.Get("token")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("-----------------------------------------------")
		fmt.Println(claims["UserId"])
	} else {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"error": nil,
		"data": gin.H{
			"token": "jiji",
		},
		"status": "Successfull",
	})
}
