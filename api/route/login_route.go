package route

import (
	"gin-jwt-gorm/MySQL"
	"gin-jwt-gorm/api/controller"
	"gin-jwt-gorm/config"
	"gin-jwt-gorm/domain/usecase"
	"gin-jwt-gorm/repository"
	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *config.Config, db MySQL.MysqlDataBase, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
