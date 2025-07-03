package main

import (
	"github.com/bigh0124/gin-gorm-redis/internal/config"
	"github.com/bigh0124/gin-gorm-redis/internal/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	r.Run(config.AppConfig.App.Port)
}
