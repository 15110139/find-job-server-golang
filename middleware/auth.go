package middleware

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/find-job-server-golang/util/constant"
	response "github.com/find-job-server-golang/util/response"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(c *gin.Context) {
	mySigningKey := []byte("AllYourBase")
	tokenString := c.Request.Header.Get("token")
	if tokenString == "" {
		response.RespondWithError(c, "TOKEN_REQUIRE", 500)
	} else {
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return mySigningKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userIdFormToken", claims["userId"].(string))
			c.Next()
		} else {
			fmt.Println(err)
			response.RespondWithError(c, constant.INVALID_TOKEN, 500)
		}
	}

}

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("userId", "123213")
		c.Writer.Header().Add("userId", "21313")
		c.Next()
	}
}
