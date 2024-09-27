package config

import "gin-jwt-gorm/MySQL"

type App struct {
	Env   *Config
	Mysql MySQL.MysqlDataBase
}

func NewApp() App {
	var app App
	app.Env = ParseConfig()
	mysqlDB, err := Connect(app.Env)
	if err != nil {
		panic(err)
	}
	app.Mysql = *mysqlDB
	return app
}
