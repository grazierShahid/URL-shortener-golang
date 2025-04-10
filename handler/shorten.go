package handler

import (
	"fmt"
	"net/http"
	"url-shortener/redis"
	"url-shortener/utils"

	"github.com/gin-gonic/gin"
	redis_v9 "github.com/redis/go-redis/v9"
)

func Shorten(c *gin.Context) {
	var requestBody struct {
		URL string `json:"url"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	shortKey := utils.GenerateRandomKey(6)

	err := redis.Client.Set(redis.Ctx, shortKey, requestBody.URL, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save to Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short_url": fmt.Sprintf("http://localhost:8080/%s", shortKey),
	})
}

func Redirect(c *gin.Context) {
	shortKey := c.Param("key")

	url, err := redis.Client.Get(redis.Ctx, shortKey).Result()
	if err == redis_v9.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch from Redis"})
		return
	}

	RecordClick(c)

	c.Redirect(http.StatusFound, url)
}
