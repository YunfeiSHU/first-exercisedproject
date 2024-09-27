package controller

import (
	"gin-jwt-gorm/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProfileController struct {
	ProfileController model.ProfileUsecase
}

func (pc *ProfileController) GetProfile(c *gin.Context) {
	userID := c.MustGet("x-user-id").(int)

	profile, err := pc.ProfileController.GetProfileByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, profile)

}
