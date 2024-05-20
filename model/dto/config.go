package dto

import "github.com/rs/zerolog"

type (
	ConfigData struct {
		DBConfig     dbConfig
		AppConfig    appConfig
		Version      string
		ClientID     string
		ClientSecret string
	}

	dbConfig struct {
		Host        string
		Port        int
		User        string
		Pass        string
		Name        string
		MaxIdle     int
		MaxConn     int
		MaxLifetime string
		MaxIdletime string
		LogLevel    int
	}

	appConfig struct {
		LogMode int
		Port    string
	}
)

// MarshalZerologObject implements zerolog.LogObjectMarshaler.
func (c *ConfigData) MarshalZerologObject(e *zerolog.Event) {
	e.Str("DB HOST", c.DBConfig.Host)
	e.Int("DB PORT", c.DBConfig.Port)
	e.Str("DB USER", c.DBConfig.User)
	e.Str("DB PASS", c.DBConfig.Pass)
	e.Str("DB NAME", c.DBConfig.Name)
	e.Int("DB MAX IDLE", c.DBConfig.MaxIdle)
	e.Int("DB MAX CONN", c.DBConfig.MaxConn)
	e.Str("DB MAX LIFETIME", c.DBConfig.MaxLifetime)
	e.Int("DB LOG LEVEL", c.AppConfig.LogMode)
	e.Str("APP PORT", c.AppConfig.Port)
	e.Int("APP LOG MODE", c.AppConfig.LogMode)
	e.Str("VERSION", c.Version)
	e.Str("CLIENT ID", c.ClientID)
	e.Str("CLIENT SECRET", c.ClientSecret)
}
