package http

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/enockjeremi/queswer/delivery/http/middlewares"
	"github.com/enockjeremi/queswer/domain"
	"github.com/enockjeremi/queswer/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
}

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PasswordInput struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=7"`
}

func NewAuthHandler(r *gin.Engine, au domain.AuthUsecase) {
	handler := &AuthHandler{
		AuthUsecase: au,
	}

	v1 := r.Group("/v1")

	v1.POST("auth/sign-in", handler.SignIn)
	v1.POST("auth/sign-up", handler.SignUp)
	v1.GET("auth/profile", middlewares.CheckAuth, handler.GetProfile)
	v1.PUT("auth/update-profile", middlewares.CheckAuth, handler.UpdateProfile)
	v1.PATCH("auth/change-password", middlewares.CheckAuth, handler.ChangePassword)
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var signIn AuthInput
	var user domain.User

	if err := c.ShouldBindBodyWithJSON(&signIn); err != nil {
		utils.BadRequestResponse(c, err)
		return
	}

	err := h.AuthUsecase.VerifyUsername(&user, signIn.Username)
	if err != nil {
		utils.ForbiddenResponse(c, "invalid username")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signIn.Password)); err != nil {
		utils.ForbiddenResponse(c, "invalid password")
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("jWT_SECRET")))
	if err != nil {
		utils.NotFoundResponse(c, "failed to generate token")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   token,
	})

}
func (h *AuthHandler) SignUp(c *gin.Context) {
	var auth domain.User

	if err := c.ShouldBindBodyWithJSON(&auth); err != nil {
		utils.BadRequestResponse(c, err)
		return
	}
	if err := h.AuthUsecase.VerifyCredentials(&auth); err == nil {
		utils.NotFoundResponse(c, "user or email already exists")
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.NotFoundResponse(c, err.Error())
		return
	}

	profile := auth.Profile
	err = h.AuthUsecase.CreateProfile(&profile)
	if err != nil {
		utils.NotFoundResponse(c, "something wrong! could not register user")
		return
	}

	user := domain.User{
		Username:  auth.Username,
		Email:     auth.Email,
		Password:  string(passwordHash),
		ProfileID: profile.ID,
		Profile:   profile,
	}

	err = h.AuthUsecase.RegisterUser(&user)
	if err != nil {
		utils.NotFoundResponse(c, "something wrong! could not register user")
		return
	} else {
		c.JSON(http.StatusCreated, user)
	}
}
func (h *AuthHandler) GetProfile(c *gin.Context) {
	user, _ := c.Get("currentUser")
	c.JSON(200, gin.H{
		"success": true,
		"user":    user,
	})
}
func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	var profile domain.Profile
	var user domain.User
	currentUser, _ := c.Get("currentUser")
	profileId := currentUser.(domain.User).ProfileID

	err := h.AuthUsecase.GetProfile(&profile, profileId)
	if err != nil {
		utils.NotFoundResponse(c, "profile not found")
		return
	}

	if err = c.ShouldBindBodyWithJSON(&profile); err != nil {
		utils.BadRequestResponse(c, err)
		return
	}

	err = h.AuthUsecase.UpdateProfile(&profile)
	if err != nil {
		utils.NotFoundResponse(c, fmt.Sprintf("Could not update user profile ID: %v", profileId))
		return
	}

	err = h.AuthUsecase.GetUser(&user, profile.UserID)
	if err != nil {
		utils.NotFoundResponse(c, "profile not found")
		return
	}
	c.JSON(http.StatusOK, user)

}
func (h *AuthHandler) ChangePassword(c *gin.Context) {
	var passwordInput PasswordInput
	var user domain.User
	currentUser, _ := c.Get("currentUser")
	id := currentUser.(domain.User).ID

	if err := c.ShouldBindBodyWithJSON(&passwordInput); err != nil {
		utils.BadRequestResponse(c, err)
		return
	}

	err := h.AuthUsecase.GetUser(&user, id)
	if err != nil {
		utils.NotFoundResponse(c, "user not found")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordInput.OldPassword)); err != nil {
		utils.NotFoundResponse(c, "old password not found")
		return
	}

	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(passwordInput.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.NotFoundResponse(c, err.Error())
		return
	}

	err = h.AuthUsecase.ChangePassword(&user, string(newPasswordHash))
	if err != nil {
		utils.NotFoundResponse(c, "could not change password")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": "successful password change",
		"success":  true,
	})

}
