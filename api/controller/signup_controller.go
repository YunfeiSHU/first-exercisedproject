package controller

import (
	"gin-jwt-gorm/config"
	"gin-jwt-gorm/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUpController struct {
	SignupUsecase model.SignupUsecase
	Env           *config.Config
}

func (sc *SignUpController) SignUp(c *gin.Context) {
	var request model.SignupRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error()})
		return
	}

	_, err := sc.SignupUsecase.GetUserByEmail(request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, model.ErrorResponse{
			Message: "User with this email already exists",
		})
	}

	user := model.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
	}

	err = sc.SignupUsecase.Create(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.RefreshTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	signupResponse := model.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
