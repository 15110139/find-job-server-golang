package middleware

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/testgolang/util"
)

func TokenAuthMiddleware(c *gin.Context) {
	mySigningKey := []byte("AllYourBase")

	tokenString := c.Request.Header.Get("token")
	if tokenString == "" {
		util.RespondWithError(c, "TOKEN_REQUIRE")
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return mySigningKey, nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
		c.Next()
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["foo"], claims["nbf"])
		} else {
			fmt.Println(err)
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
}
