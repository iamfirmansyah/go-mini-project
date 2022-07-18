package config

import (
	"go-pos/helper"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PORT                             string `mapstructure:"PORT"`
	MYSQL_CONNECTION                 string `mapstructure:"MYSQL_CONNECTION"`
	MONGO_DB_HOST                    string `mapstructure:"MONGO_DB_HOST"`
	MONGO_DB_PORT                    string `mapstructure:"MONGO_DB_PORT"`
	MONGO_DB_DATABASE                string `mapstructure:"MONGO_DB_DATABASE"`
	MONGO_DB_USERNAME                string `mapstructure:"MONGO_DB_USERNAME"`
	MONGO_DB_PASSWORD                string `mapstructure:"MONGO_DB_PASSWORD"`
	MONGO_DB_AUTHENTICATION_DATABASE string `mapstructure:"MONGO_DB_AUTHENTICATION_DATABASE"`
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	helper.PanicIfError(err)

	err = viper.Unmarshal(&AppConfig)
	helper.PanicIfError(err)
	log.Println("Configuration Success...")
}
