package controllers

import (
	"net/http"

	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUser(c *gin.Context) {}

func RegisterUser(c *gin.Context) {
	var auth models.User

	if err := c.ShouldBindBodyWithJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.VerifyCredentials(&auth); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user or email already exists"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: auth.Username,
		Email:    auth.Email,
		Password: string(passwordHash),
	}

	err = services.CreateUser(&user)
	if err != nil {
		NotFoundError(c, "Could not register user")
		return
	} else {
		c.JSON(http.StatusCreated, user)
	}

}

func GetOneUser(c *gin.Context) {}
func PutUser(c *gin.Context)    {}
func DeleteUser(c *gin.Context) {}
