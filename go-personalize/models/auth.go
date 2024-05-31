package models

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	//Get the cookies
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])

		}
		return []byte("bfuigfebfyuie"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//Check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//Find the user with token
		var user Register
		DB.First(&user, claims["sub"])

		if user.Id == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//Attach to req
		c.Set("user", user)

		//Continue
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
