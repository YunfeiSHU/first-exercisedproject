package controller

import (
	"gin-jwt-gorm/config"
	"gin-jwt-gorm/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//与model层的LoginUsecase对接：将实现了model.Login接口的类model.LoginUsecase 嵌套

type LoginController struct {
	LoginUsecase model.LoginUsecase
	Env          *config.Config
}

func (lc *LoginController) Login(c *gin.Context) {
	//绑定登录时信息
	var request model.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	//验证邮箱
	user, err := lc.LoginUsecase.GetUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if request.Password != user.Password {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Message: "password error",
		})
		return
	}

	//创建token
	accessToken, err := lc.LoginUsecase.CreateAccessToken(user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	//返回token
	loginResponse := model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, loginResponse)
}
