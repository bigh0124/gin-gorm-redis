package handler

import (
	"errors"
	"net/http"

	"github.com/bigh0124/gin-gorm-redis/internal/config"
	"github.com/bigh0124/gin-gorm-redis/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
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

func GetArticles(c *gin.Context) {
	db := config.GetDB()

	var articles []model.Article

	if err := db.Find(&articles).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}

func GetArticleByID(c *gin.Context) {
	db := config.GetDB()

	var article model.Article

	id, ok := c.Params.Get("ID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "url not allow",
		})
	}

	if err := db.Where("ID = ?", id).First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}

func LikeArticle(c *gin.Context) {
	ID, ok := c.Params.Get("ID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid url",
		})
		return
	}

	rdb := config.GetRedis()

	likeKey := "article:" + ID + ":like"
	if err := rdb.Incr(c.Request.Context(), likeKey).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully like article",
	})
}

func GetArticleLikes(c *gin.Context) {
	ID, ok := c.Params.Get("ID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid url",
		})
		return
	}

	rdb := config.GetRedis()

	likeKey := "article:" + ID + ":like"
	likes, err := rdb.Get(c.Request.Context(), likeKey).Result()
	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"likes": likes,
	})
}
