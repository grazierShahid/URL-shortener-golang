package main

import (
	"url-shortener/database"
	"url-shortener/handler"
	"url-shortener/redis"

	"github.com/gin-gonic/gin"
)

func init() {
	redis.InitRedis()
	database.InitDB()
	handler.InitIP2Location()
}

func main() {
	r := gin.Default()

	r.POST("/api/shorten", handler.Shorten)

	r.GET("/:key", handler.Redirect)

	r.GET("/api/analytics/:key", handler.GetURLAnalytics)

	r.GET("/api/stats", handler.GetAllURLStats)

	r.Run(":8080")
}
