package config

import (
	"errors"
	"finpro-fenlie/entity/dto"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func InitEnv() (*dto.ConfigData, error) {
	var configData dto.ConfigData

	// load .env file from root
	if err := godotenv.Load(); err != nil {
		return &configData, err
	}

	if version := os.Getenv("VERSION"); version != "" {
		configData.Version = version
	}

	if port := os.Getenv("PORT"); port != "" {
		configData.AppConfig.Port = port
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbMaxIdle := os.Getenv("MAX_IDLE")
	dbMaxConn := os.Getenv("MAX_CONN")
	dbMaxLifetime := os.Getenv("MAX_LIFE_TIME")
	logMode := os.Getenv("LOG_MODE")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPass == "" || dbName == "" ||
		dbMaxIdle == "" || dbMaxConn == "" || dbMaxLifetime == "" || logMode == "" {
		return &configData, errors.New("DB Config is not set")
	}

	var err error
	configData.DBConfig.MaxConn, err = strconv.Atoi(dbMaxConn)
	if err != nil {
		return &configData, err
	}

	configData.DBConfig.MaxIdle, err = strconv.Atoi(dbMaxIdle)
	if err != nil {
		return &configData, err
	}

	configData.DBConfig.Port, err = strconv.Atoi(dbPort)
	if err != nil {
		return &configData, err
	}

	configData.DBConfig.Host = dbHost
	configData.DBConfig.User = dbUser
	configData.DBConfig.Pass = dbPass
	configData.DBConfig.Name = dbName
	configData.DBConfig.MaxLifetime = dbMaxLifetime

	configData.DBConfig.LogMode, err = strconv.Atoi(logMode)
	if err != nil {
		return &configData, err
	}

	return &configData, nil
}
