package route

import (
	"gin-jwt-gorm/MySQL"
	"gin-jwt-gorm/api/controller"
	"gin-jwt-gorm/config"
	"gin-jwt-gorm/domain/usecase"
	"gin-jwt-gorm/repository"
	"github.com/gin-gonic/gin"
)

func NewProfileRouter(env *config.Config, db MySQL.MysqlDataBase, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	pc := &controller.ProfileController{
		ProfileController: usecase.NewProfileUsecase(ur),
	}
	group.GET("/profile", pc.GetProfile)
}
