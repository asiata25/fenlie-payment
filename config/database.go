package config

import (
	"finpro-fenlie/model/dto"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type zerologWriter struct {
	Logger zerolog.Logger
}

func (l *zerologWriter) Printf(message string, args ...interface{}) {
	l.Logger.Info().Msgf(message, args...)
}

func ConnectDB(in *dto.ConfigData, log zerolog.Logger) (*gorm.DB, error) {
	log.Info().Msg("Trying to connect to the database . . .")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", in.DBConfig.Host, in.DBConfig.User, in.DBConfig.Pass, in.DBConfig.Name, in.DBConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // Disable by default transaction
		Logger: logger.New(&zerologWriter{Logger: log}, logger.Config{ // overwrite the default logger by using zerolog
			LogLevel: logger.LogLevel(in.DBConfig.LogLevel),
			Colorful: true,
		}),
	})

	// GORM automatically ping database after initialized to check database availability

	if err != nil {
		err = errors.New(err.Error())
		log.Error().Err(err).Msg("Failed to open database connection")
		return nil, err
	}

	log.Info().Msg("Successfully connect to the database")
	return db, nil
}
