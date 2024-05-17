package app

import (
	"finpro-fenlie/config"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func RunService() {
	// this is where the service will run

	// set project timezone
	time.Local = time.FixedZone("Asia/Jakarta", 7*60*60)

	// set global logger with zerolog
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().Caller().Logger()

	// setup config file
	configData, err := config.InitEnv()
	if err != nil {
		log.Error().Err(err).Msg("Failed to load config .env")
		return
	}
	log.Info().Object("config", configData).Msg("Success load .env")

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

	// set gin middleware for logging
	r.Use(logger.SetLogger(logger.WithLogger(func(ctx *gin.Context, l zerolog.Logger) zerolog.Logger {
		return l.Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().Logger()
	})))

	// set gin middleware for panic hanlder
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	version := configData.Version
	log.Info().Msgf("Service running version: %s", version)
	port := configData.AppConfig.Port
	err = r.Run(port)
	if err != nil {
		log.Panic().Err(err).Msgf("Failed to run service on port %s", port)
	}
}
