package auth

import (
	"fmt"
	"net/http"

	"github.com/enockjeremi/queswer/libs"
	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	user, _ := c.Get("currentUser")
	c.JSON(200, gin.H{
		"success": true,
		"user":    user,
	})
}

func UpdateProfile(c *gin.Context) {
	var profile models.Profile
	var user models.User
	currentUser, _ := c.Get("currentUser")
	profileId := currentUser.(models.User).ProfileID

	err := services.GetProfile(&profile, profileId)
	if err != nil {
		libs.NotFoundResponse(c, "profile not found")
		return
	}

	if err = c.ShouldBindBodyWithJSON(&profile); err != nil {
		libs.BadRequestResponse(c, err)
		return
	}

	err = services.UpdateProfile(&profile)
	if err != nil {
		libs.NotFoundResponse(c, fmt.Sprintf("Could not update user profile ID: %v", profileId))
		return
	}

	err = services.GetUser(&user, profile.UserID)
	if err != nil {
		libs.NotFoundResponse(c, "profile not found")
		return
	}
	c.JSON(http.StatusOK, user)

}
