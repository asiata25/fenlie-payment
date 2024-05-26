package config

import (
	"errors"
	"finpro-fenlie/model/dto"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func InitEnv() (*dto.ConfigData, error) {
	var configData dto.ConfigData

	// load .env file from root
	mode := os.Getenv("MODE")
	if mode != "" {
		if err := godotenv.Load(mode + ".env"); err != nil {
			return &configData, err
		}
	} else {
		if err := godotenv.Load(); err != nil {
			return &configData, err
		}
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
	dbMaxIdletime := os.Getenv("MAX_IDLE_TIME")
	dbLogLevel := os.Getenv("DB_LOG_LEVEL")
	logMode := os.Getenv("LOG_MODE")
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	brickId := os.Getenv("BRICK_ID")
	brickSecret := os.Getenv("BRICK_SECRET")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPass == "" || dbName == "" ||
		dbMaxIdle == "" || dbMaxConn == "" || dbMaxLifetime == "" || logMode == "" || dbLogLevel == "" || clientId == "" || clientSecret == "" || brickId == "" || brickSecret == "" {
		return &configData, errors.New("DB Config is not properly set")
	}

	var err error
	configData.DBConfig.MaxConn, err = strconv.Atoi(dbMaxConn)
	if err != nil {
		return &configData, err
	}

	configData.DBConfig.LogLevel, err = strconv.Atoi(dbLogLevel)
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
	configData.DBConfig.MaxIdletime = dbMaxIdletime

	configData.AppConfig.LogMode, err = strconv.Atoi(logMode)
	if err != nil {
		return &configData, err
	}

	configData.ClientID = clientId
	configData.ClientSecret = clientSecret
	configData.BrickID = brickId
	configData.BrickSecret = brickSecret

	return &configData, nil
}
