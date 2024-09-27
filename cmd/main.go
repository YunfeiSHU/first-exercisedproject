package main

import (
	"gin-jwt-gorm/api/route"
	"gin-jwt-gorm/config"
	"github.com/gin-gonic/gin"
)

func main() {
	app := config.NewApp()

	env := app.Env
	db := app.Mysql
	router := gin.Default()

	route.Setup(env, db, router)
	_ = router.Run(":8080")
}
