package controller

import (
	"gin-jwt-gorm/config"
	"gin-jwt-gorm/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RefreshTokenController struct {
	RefreshTokenUsecase model.RefreshTokenUsecase
	Env                 *config.Config
}

func (rtc *RefreshTokenController) RefreshToken(c *gin.Context) {
	var request model.RefreshTokenRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	idstring, err := rtc.RefreshTokenUsecase.ExtractIDFromToken(request.RefreshToken, rtc.Env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(idstring)

	user, err := rtc.RefreshTokenUsecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	accessToken, err := rtc.RefreshTokenUsecase.CreateAccessToken(user, rtc.Env.RefreshTokenSecret, rtc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	refreshToken, err := rtc.RefreshTokenUsecase.CreateRefreshToken(user, rtc.Env.RefreshTokenSecret, rtc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response := model.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, response)
}
