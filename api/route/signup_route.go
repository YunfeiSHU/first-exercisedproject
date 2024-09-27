package route

import (
	"gin-jwt-gorm/MySQL"
	"gin-jwt-gorm/api/controller"
	"gin-jwt-gorm/config"
	"gin-jwt-gorm/domain/usecase"
	"gin-jwt-gorm/repository"
	"github.com/gin-gonic/gin"
)

func NewSignupRouter(env *config.Config, db MySQL.MysqlDataBase, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	sc := controller.SignUpController{
		SignupUsecase: usecase.NewSignupUsecase(ur),
		Env:           env,
	}
	group.POST("/signup", sc.SignUp)
}
