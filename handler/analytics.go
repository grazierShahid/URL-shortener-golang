package handler

import (
	"log"
	"net/http"
	"time"
	"url-shortener/database"

	"github.com/gin-gonic/gin"
	"github.com/ip2location/ip2location-go/v9"
)

var ip2l *ip2location.DB

func InitIP2Location() {
	var err error
	ip2l, err = ip2location.OpenDB("./IP2LOCATION-LITE-DB5.BIN")
	if err != nil {
		log.Printf("Warning: IP2Location database not found. Geolocation features will be disabled: %v", err)
		ip2l = nil
	} else {
		log.Println("IP2Location database loaded successfully")
	}
}

func RecordClick(c *gin.Context) {
	shortKey := c.Param("key")

	ipAddress := c.ClientIP()

	var country, city, region string

	if ip2l != nil {
		results, err := ip2l.Get_all(ipAddress)
		if err == nil {
			country = results.Country_long
			city = results.City
			region = results.Region
		}
	}

	_, err := database.DB.Exec(
		"INSERT INTO url_clicks (short_key, click_time, ip_address, country, city, region) VALUES ($1, $2, $3, $4, $5, $6)",
		shortKey,
		time.Now(),
		ipAddress,
		country,
		city,
		region,
	)

	if err != nil {
		c.Error(err)
	}
}

func GetURLAnalytics(c *gin.Context) {
	shortKey := c.Param("key")

	var totalClicks int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM url_clicks WHERE short_key = $1", shortKey).Scan(&totalClicks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch analytics data"})
		return
	}

	rows, err := database.DB.Query(
		"SELECT country, COUNT(*) as count FROM url_clicks WHERE short_key = $1 GROUP BY country ORDER BY count DESC",
		shortKey,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch country analytics"})
		return
	}
	defer rows.Close()

	countryStats := []gin.H{}
	for rows.Next() {
		var country string
		var count int
		if err := rows.Scan(&country, &count); err != nil {
			continue
		}
		countryStats = append(countryStats, gin.H{"country": country, "count": count})
	}

	timeRows, err := database.DB.Query(
		"SELECT DATE(click_time) as date, COUNT(*) as count FROM url_clicks "+
			"WHERE short_key = $1 AND click_time > NOW() - INTERVAL '30 days' "+
			"GROUP BY date ORDER BY date",
		shortKey,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch time analytics"})
		return
	}
	defer timeRows.Close()

	timeStats := []gin.H{}
	for timeRows.Next() {
		var date time.Time
		var count int
		if err := timeRows.Scan(&date, &count); err != nil {
			continue
		}
		timeStats = append(timeStats, gin.H{"date": date.Format("2006-01-02"), "count": count})
	}

	c.JSON(http.StatusOK, gin.H{
		"short_key":     shortKey,
		"total_clicks":  totalClicks,
		"country_stats": countryStats,
		"time_stats":    timeStats,
	})
}
