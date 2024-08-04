package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/enockjeremi/queswer/config"
	"github.com/enockjeremi/queswer/domain"
	"github.com/enockjeremi/queswer/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CheckAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		utils.UnauthorizedResponse(c, "authorization header is missing")
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
		utils.UnauthorizedResponse(c, "invalid or expired token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		utils.UnauthorizedResponse(c, "invalid token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		utils.UnauthorizedResponse(c, "token expired")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userID := fmt.Sprintf("%v", (claims["id"]))
	var user domain.User
	if err := config.DB.Where("id = ?", userID).Preload("Profile").First(&user).Error; err != nil {
		utils.UnauthorizedResponse(c, "could not authenticate user")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("currentUser", user)
	c.Next()
}
