package main

import (
	"github.com/bigh0124/gin-gorm-redis/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(config.AppConfig.App.Port)
}
