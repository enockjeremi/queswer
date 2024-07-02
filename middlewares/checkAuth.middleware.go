package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/enockjeremi/queswer/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		utils.ErrorHandling(c, http.StatusUnauthorized, "Authorization header is missing")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	authToken := strings.Split(authHeader, " ")

	tokenString := authToken[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("jWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		utils.ErrorHandling(c, http.StatusUnauthorized, "invalid expired token")

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		utils.ErrorHandling(c, http.StatusUnauthorized, "invalid token")
		c.Abort()
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		utils.ErrorHandling(c, http.StatusUnauthorized, "token expired")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User
	err = services.GetUser(&user, fmt.Sprintf("%v", (claims["id"])))
	if err != nil {
		utils.ErrorHandling(c, http.StatusUnauthorized, "invalid token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("currentUser", user)
	c.Next()
}
