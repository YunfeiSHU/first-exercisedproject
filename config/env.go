package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host                   string `yaml:"host"`
	Port                   string `yaml:"port"`
	User                   string `yaml:"user"`
	Password               string `yaml:"password"`
	DBName                 string `yaml:"dbname"`
	AccessTokenExpiryHour  int    `toml:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `toml:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `toml:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `toml:"REFRESH_TOKEN_SECRET"`
}

func ParseConfig() (dbconfig *Config) {
	config := viper.New()
	config.SetConfigFile("./config.toml")
	config.SetConfigType("toml")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := config.Unmarshal(&dbconfig); err != nil {
		panic(err)
	}
	return dbconfig
}
