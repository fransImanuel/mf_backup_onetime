package dto

import "github.com/gin-contrib/cors"

type AppConfig struct {
	Debug        bool
	Host         string
	Port         string
	UseTLS       bool
	CertFilePath string
	KeyFilePath  string
	Cors         *cors.Config
}

type DatabaseConfig struct {
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	Dialect    string
	TimeZone   string
	AuthSource string
	SSLMode    string
}

type RedisConfig struct {
	Host     string
	Password string
	DB       int
}

type SMTPConfig struct {
	Host     string
	Port     int
	Email    string
	Password string
	Name     string
}

type OneSignalConfig struct {
	AppId      string
	RestApiKey string
}

type NewRelicConfig struct {
	AppName string
	License string
}

type MinioConfig struct {
	Host          string
	Location      string
	AccessKey     string
	SecretKey     string
	SSL           bool
	ReplaceDomain string
}

type ConfigValue struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type ConfigResponse struct {
	Value     string `json:"value" `
	UpdatedAt string `json:"updated_at" `
}
