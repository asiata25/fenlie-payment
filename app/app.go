package app

import (
	"finpro-fenlie/config"
	"finpro-fenlie/pkg/validation"
	"finpro-fenlie/router"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"gorm.io/gorm"
)

func initializeDomainModule(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")
	v1 := api.Group("/v1")

	router.InitRouter(v1, db)

}

func RunService() {
	// this is where the service will run

	// set project timezone
	time.Local = time.FixedZone("Asia/Jakarta", 7*60*60)

	// set global logger with zerolog
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	// Setup output to console
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}

	// Setup output to file
	// logPath := "log"
	// FIXME: log file will be created only once when the first time app running
	// logFileName := time.Now().Format("2006-01-02") + ".log"
	// logFilePath := filepath.Join(logPath, logFileName)

	// Create log directory if not exists
	// if _, err := os.Stat(logPath); os.IsNotExist(err) {
	// 	os.Mkdir(logPath, os.ModePerm)
	// }

	// Open log file
	// logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("Failed to open log file")
	// }
	// defer logFile.Close()

	// Setup output to file
	// fileWriter := zerolog.ConsoleWriter{Out: logFile}

	// Combine console and file output
	multi := zerolog.MultiLevelWriter(consoleWriter)

	// Set logger output
	log.Logger = log.Output(multi).With().Stack().Caller().Logger()

	// Filter out messages with level lower than error
	// errorLogger := zerolog.New(multi).Level(zerolog.ErrorLevel)

	// Set logger output to file for error level logs
	// FIXME: log.Logger above will be overwrite, so in database will only show error log
	// log.Logger = errorLogger

	// setup config file
	configData, err := config.InitEnv()
	if err != nil {
		log.Error().Err(err).Msg("failed to load config .env")
		return
	}
	log.Info().Object("config", configData).Msg("Success load " + os.Getenv("MODE") + ".env")

	// setup database
	conn, err := config.ConnectDB(configData, log.Logger)
	if err != nil {
		log.Error().Msgf("RunService.ConnectDB.err : %s", err.Error())
		return
	}

	durationConn, err := time.ParseDuration(configData.DBConfig.MaxLifetime)
	if err != nil {
		log.Error().Msgf("RunService.ConnDuration.err :%s", err.Error())
		return
	}
	durationIdle, err := time.ParseDuration(configData.DBConfig.MaxIdletime)
	if err != nil {
		log.Error().Msgf("RunService.IdleDuration.err :%s", err.Error())
		return
	}

	sqlDB, err := conn.DB()
	if err != nil {
		log.Error().Err(err).Msgf("RunService.sqlDB")
		return
	}

	sqlDB.SetConnMaxLifetime(durationConn)
	sqlDB.SetMaxIdleConns(configData.DBConfig.MaxIdle)
	sqlDB.SetConnMaxIdleTime(durationIdle)
	sqlDB.SetMaxOpenConns(configData.DBConfig.MaxConn)

	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close DB")
		}
	}()

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
		return l.Level(zerolog.Level(configData.AppConfig.LogMode)).Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().Logger()
	})))

	// set gin middleware for panic hanlder
	r.Use(gin.Recovery())

	// add custom validation for password
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("isPassword", validation.IsPassword)
	}

	// setup global endpoint
	initializeDomainModule(r, conn)

	version := configData.Version
	log.Info().Msgf("service running version: %s", version)
	port := configData.AppConfig.Port
	err = r.Run(port)
	if err != nil {
		log.Panic().Err(err).Msgf("failed to run service on port %s", port)
	}
}
