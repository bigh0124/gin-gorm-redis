package router

import (
	"github.com/bigh0124/gin-gorm-redis/internal/handler"
	"github.com/bigh0124/gin-gorm-redis/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}

	article := r.Group("/api/article")
	article.Use(middleware.Authenticate())
	{
		article.POST("/", handler.CreateArticleHandler)
		article.GET("/", handler.GetArticles)
	}

	return r
}
