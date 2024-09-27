package route

import (
	"gin-jwt-gorm/MySQL"
	"gin-jwt-gorm/api/controller"
	"gin-jwt-gorm/config"
	"gin-jwt-gorm/domain/usecase"
	"gin-jwt-gorm/repository"
	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(env *config.Config, db MySQL.MysqlDataBase, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
