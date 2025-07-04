package handler

import (
	"net/http"

	"github.com/bigh0124/gin-gorm-redis/internal/config"
	"github.com/bigh0124/gin-gorm-redis/internal/model"
	"github.com/gin-gonic/gin"
)

func CreateArticleHandler(c *gin.Context) {
	var article model.Article

	if err := c.ShouldBind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db := config.GetDB()

	if err := db.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}
