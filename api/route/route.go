package route

import (
	"gin-jwt-gorm/MySQL"
	"gin-jwt-gorm/api/controller"
	"gin-jwt-gorm/api/middleware"
	"gin-jwt-gorm/config"
	"github.com/gin-gonic/gin"
)

func Setup(env *config.Config, db MySQL.MysqlDataBase, gin *gin.Engine) {
	//public api
	publicRouter := gin.Group("api/v1")
	NewSignupRouter(env, db, publicRouter)
	NewLoginRouter(env, db, publicRouter)
	NewRefreshTokenRouter(env, db, publicRouter)
	//middleware to vertify accesstoken
	protectedRouter := gin.Group("api/v1/protected")
	protectedRouter.Use(middleware.JwtAuthMidddleware(env.AccessTokenSecret))
	//private api
	NewProfileRouter(env, db, protectedRouter)

	//test api
	testRouter := gin.Group("api/test")
	testRouter.GET("/ping", controller.Ping)
}
