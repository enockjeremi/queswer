package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/enockjeremi/queswer/config"
	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignIn(c *gin.Context) {
	var signIn SignInInput
	var auth models.User

	if err := c.ShouldBindBodyWithJSON(&signIn); err != nil {
		ErrorHandling(c, http.StatusBadRequest, err.Error())
		return
	}

	config.DB.Where("username = ?", signIn.Username).Find(&auth)
	if auth.ID == 0 {
		ErrorHandling(c, http.StatusForbidden, "invalid user or password")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(signIn.Password)); err != nil {
		ErrorHandling(c, http.StatusForbidden, "invalid user or password")
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id: ": auth.ID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("jWT_SECRET")))
	if err != nil {
		ErrorHandling(c, http.StatusBadRequest, "failed to generate token")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   token,
	})

}

func SignUp(c *gin.Context) {
	var auth models.User

	if err := c.ShouldBindBodyWithJSON(&auth); err != nil {
		ErrorHandling(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := services.VerifyCredentials(&auth); err == nil {
		ErrorHandling(c, http.StatusBadRequest, "user or email already exists")
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		ErrorHandling(c, http.StatusNotFound, err.Error())
		return
	}

	user := models.User{
		Username: auth.Username,
		Email:    auth.Email,
		Password: string(passwordHash),
	}

	err = services.CreateUser(&user)
	if err != nil {
		ErrorHandling(c, 404, "Could not register user")
		return
	} else {
		c.JSON(http.StatusCreated, user)
	}

}
