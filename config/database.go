package config

import (
	"fmt"
	"gin-jwt-gorm/MySQL"

	"gin-jwt-gorm/domain/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(dbConfig *Config) (*MySQL.MysqlDataBase, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.User, dbConfig.Password,
		dbConfig.Host, dbConfig.Port, dbConfig.DBName)

	//如果你要在方法中，赋值全局DB；确保err在参数列表中，就定义了
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		return nil, err
	}

	mysqldb := &MySQL.MysqlDataBase{DB: db}

	_ = db.AutoMigrate(&model.User{})

	return mysqldb, nil
}
