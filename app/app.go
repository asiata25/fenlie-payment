package app

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunService() {
	// this is where the service will run

	// set project timezone
	time.Local = time.FixedZone("Asia/Jakarta", 7*60*60)

	// initialize gin for router framefork
	r := gin.New()

	// set gin middleware for cors handler
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: false,
		AllowOrigins:    []string{"*"},
		AllowMethods:    []string{"POST", "DELETE", "GET", "OPTIONS", "PUT"},
		AllowHeaders: []string{
			"Origini", "Content-Type",
			"Authorization",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           120 * time.Second,
	}))

	// set gin middleware for panic hanlder
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
