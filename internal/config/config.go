package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig = &Config{}

	err := viper.Unmarshal(AppConfig)
	if err != nil {
		log.Fatalf("Error unmarshal config: %v", err)
	}

	err = InitDB()
	if err != nil {
		log.Fatalf("Error db connection: %v", err)
	}

	err = InitRedis()
	if err != nil {
		log.Fatalf("Error redis connection: %v", err)
	}
}
