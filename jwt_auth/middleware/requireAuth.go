package middleware

import (
	"fmt"
	"jwtAuth/initial"
	model "jwtAuth/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	// Get tocken from cookie
	tokenStr, err := c.Cookie("Auth")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Validate
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Check exp
	if time.Now().Unix() > claims["exp"].(int64) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	
	// Get user
	var user model.User
	initial.DB.First(&user, claims["sub"])
	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	
	// Attach user to request
	c.Set("user", user)
	
	c.Next()
}
