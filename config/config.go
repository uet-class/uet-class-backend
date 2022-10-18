package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(environment string) {
	config = viper.New()
	config.SetConfigName(environment)
	config.SetConfigType("yaml")
	config.AddConfigPath("config/")

	if err := config.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func GetConfig() *viper.Viper {
	return config
}
