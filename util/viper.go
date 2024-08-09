package util

import (
	"mf_backup_onetime/constant"
	"mf_backup_onetime/dto"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig(isDevelopment *bool) {
	log.Info("InitConfig() - starting...")

	if *isDevelopment {
		viper.SetConfigFile(constant.CONFIG_PATH_DEVELOPMENT)
	} else {
		viper.SetConfigFile(constant.CONFIG_PATH)
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("InitConfig() - error: ", err)
	}

	if viper.GetBool(constant.CONFIG_PATH_APP_DEBUG) {
		log.Info("Service run on DEBUG mode.")
	}
	log.Info("InitConfig() - finished.")
}

func GetDatabaseConfig() *dto.DatabaseConfig {
	log.Info("GetDatabaseConfig() - starting...")
	dbHost := viper.GetString(constant.CONFIG_PATH_DB_HOST)
	dbPort := viper.GetString(constant.CONFIG_PATH_DB_PORT)
	dbUsername := viper.GetString(constant.CONFIG_PATH_DB_USERNAME)
	dbPassword := viper.GetString(constant.CONFIG_PATH_DB_PASSWORD)
	dbName := viper.GetString(constant.CONFIG_PATH_DB_DATABASE)
	dbDialect := viper.GetString(constant.CONFIG_PATH_DB_DIALECT)
	dbSSLMode := viper.GetString(constant.CONFIG_PATH_DB_SSL_MODE)
	dbTimeZone := viper.GetString(constant.CONFIG_PATH_APP_TIME_ZONE)

	if dbHost == "" {
		dbHost = constant.DEFAULT_DB_HOST
	}

	if dbPort == "" {
		dbPort = constant.DEFAULT_DB_PORT
	}

	if dbUsername == "" {
		dbUsername = constant.DEFAULT_DB_USERNAME
	}

	if dbPassword == "" {
		dbPassword = constant.DEFAULT_DB_PASSWORD
	}

	if dbName == "" {
		dbName = constant.DEFAULT_DB_DATABASE
	}

	if dbDialect == "" {
		dbDialect = constant.DEFAULT_DB_DIALECT
	}

	if dbSSLMode == "" {
		dbSSLMode = constant.DEFAULT_DB_SSL_MODE
	}

	if dbTimeZone == "" {
		dbTimeZone = constant.DEFAULT_TIME_ZONE
	}

	config := &dto.DatabaseConfig{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUsername,
		Password: dbPassword,
		Database: dbName,
		Dialect:  dbDialect,
		SSLMode:  dbSSLMode,
		TimeZone: dbTimeZone,
	}

	log.Info("GetDatabaseConfig() - finished.")
	return config
}

func GetDatabaseMongoConfig() *dto.DatabaseConfig {
	log.Info("GetDatabaseConfig() - starting...")
	dbHost := viper.GetString(constant.CONFIG_PATH_DB_HOST_MONGO)
	dbPort := viper.GetString(constant.CONFIG_PATH_DB_PORT_MONGO)
	dbUsername := viper.GetString(constant.CONFIG_PATH_DB_USERNAME_MONGO)
	dbPassword := viper.GetString(constant.CONFIG_PATH_DB_PASSWORD_MONGO)
	dbName := viper.GetString(constant.CONFIG_PATH_DB_DATABASE_MONGO)
	dbSSLMode := viper.GetString(constant.CONFIG_PATH_DB_SSL_MODE_MONGO)
	dbAuthSource := viper.GetString(constant.CONFIG_PATH_DB_AUTHSOURCE_MONGO)
	dbTimeZone := viper.GetString(constant.CONFIG_PATH_APP_TIME_ZONE)

	config := &dto.DatabaseConfig{
		Host:       dbHost,
		Port:       dbPort,
		Username:   dbUsername,
		Password:   dbPassword,
		Database:   dbName,
		SSLMode:    dbSSLMode,
		AuthSource: dbAuthSource,
		TimeZone:   dbTimeZone,
	}

	log.Info("GetDatabaseConfig() - finished.")
	return config
}

func GetAppConfig() *dto.AppConfig {
	log.Info("GetAppConfig() - starting...")
	appDebug := viper.GetBool(constant.CONFIG_PATH_APP_DEBUG)
	appHost := viper.GetString(constant.CONFIG_PATH_APP_HOST)
	appPort := viper.GetString(constant.CONFIG_PATH_APP_PORT)

	appUseTLS := viper.GetBool(constant.CONFIG_PATH_APP_USE_TLS)
	appCertFilePath := viper.GetString(constant.CONFIG_PATH_APP_CERT_FILE_PATH)
	appKeyFilePath := viper.GetString(constant.CONFIG_PATH_APP_KEY_FILE_PATH)

	appCors := cors.DefaultConfig()
	appCors.AllowAllOrigins = true

	if !appDebug {
		corsAllowOriginS := viper.GetString(constant.CONFIG_PATH_ALLOW_ORIGIN)
		corsAllowOrigin := strings.Split(corsAllowOriginS, ",")
		corsAllowMethodS := viper.GetString(constant.CONFIG_PATH_ALLOW_METHOD)
		corsAllowMethod := strings.Split(corsAllowMethodS, ",")
		corsAllowHeaderS := viper.GetString(constant.CONFIG_PATH_ALLOW_HEADER)
		corsAllowHeader := strings.Split(corsAllowHeaderS, ",")
		corsExposeHeaderS := viper.GetString(constant.CONFIG_PATH_EXPOSE_HEADER)
		corsExposeHeader := strings.Split(corsExposeHeaderS, ",")
		corsAllowCredential := viper.GetBool(constant.CONFIG_PATH_ALLOW_CREDENTIAL)
		corsAllowWildcard := viper.GetBool(constant.CONFIG_PATH_ALLOW_WILDCARD)
		corsAllowBrowserExtension := viper.GetBool(constant.CONFIG_PATH_ALLOW_BROWSER_EXTENSION)
		corsAllowWebSocket := viper.GetBool(constant.CONFIG_PATH_ALLOW_WEB_SOCKET)
		corsAllowFile := viper.GetBool(constant.CONFIG_PATH_ALLOW_FILE)
		maxAgeI := viper.GetDuration(constant.CONFIG_PATH_MAX_AGE)
		maxAge := maxAgeI * time.Hour

		appCors = cors.Config{
			AllowAllOrigins:        false,
			AllowOrigins:           corsAllowOrigin,
			AllowMethods:           corsAllowMethod,
			AllowHeaders:           corsAllowHeader,
			ExposeHeaders:          corsExposeHeader,
			AllowCredentials:       corsAllowCredential,
			AllowBrowserExtensions: corsAllowBrowserExtension,
			AllowWildcard:          corsAllowWildcard,
			AllowWebSockets:        corsAllowWebSocket,
			AllowFiles:             corsAllowFile,
			MaxAge:                 maxAge,
		}
	}

	if appHost == "" {
		appHost = constant.DEFAULT_APP_HOST
	}

	if appPort == "" {
		appPort = constant.DEFAULT_APP_PORT
	}

	log.Info("GetAppConfig() - finished.")

	return &dto.AppConfig{
		Debug:        appDebug,
		Host:         appHost,
		Port:         appPort,
		UseTLS:       appUseTLS,
		CertFilePath: appCertFilePath,
		KeyFilePath:  appKeyFilePath,
		Cors:         &appCors,
	}
}

func GetRedisConfig() *dto.RedisConfig {
	log.Info("GetRedisConfig() - starting...")
	redisHost := viper.GetString(constant.CONFIG_PATH_REDIS_HOST)
	redisPassword := viper.GetString(constant.CONFIG_PATH_REDIS_PASSWORD)
	redisDB := viper.GetInt(constant.CONFIG_PATH_REDIS_DB)

	if redisHost == "" {
		redisHost = constant.DEFAULT_REDIS_HOST
	}

	if redisPassword == "" {
		redisPassword = constant.DEFAULT_DB_PASSWORD
	}

	if redisDB < 1 {
		redisDB = constant.DEFAULT_REDIS_DB
	}

	config := &dto.RedisConfig{
		Host:     redisHost,
		Password: redisPassword,
		DB:       redisDB,
	}

	log.Info("GetRedisConfig() - finished.")
	return config
}

func GetSMTPConfig() *dto.SMTPConfig {
	log.Info("GetSMTPConfig() - starting...")
	smtpHost := viper.GetString(constant.CONFIG_PATH_SMTP_HOST)
	smtpPort := viper.GetInt(constant.CONFIG_PATH_SMTP_PORT)
	smtpEmail := viper.GetString(constant.CONFIG_PATH_SMTP_EMAIL)
	smtpPassword := viper.GetString(constant.CONFIG_PATH_SMTP_PASSWORD)
	smtpName := viper.GetString(constant.CONFIG_PATH_SMTP_NAME)

	if smtpHost == "" {
		smtpHost = constant.DEFAULT_SMTP_HOST
	}

	if smtpPort < 1 {
		smtpPort = constant.DEFAULT_SMTP_PORT
	}

	if smtpEmail == "" {
		smtpEmail = constant.DEFAULT_SMTP_EMAIL
	}

	if smtpPassword == "" {
		smtpPassword = constant.DEFAULT_SMTP_PASSWORD
	}

	if smtpName == "" {
		smtpName = constant.DEFAULT_SMTP_NAME
	}

	config := &dto.SMTPConfig{
		Host:     smtpHost,
		Port:     smtpPort,
		Email:    smtpEmail,
		Password: smtpPassword,
		Name:     smtpName,
	}

	log.Info("GetSMTPConfig() - finished.")
	return config
}

func GetOneSignalConfig() *dto.OneSignalConfig {
	log.Info("GetOneSignalConfig() - starting...")
	osAppId := viper.GetString(constant.CONFIG_PATH_ONE_SIGNAL_APP_ID)
	osRestApiKey := viper.GetString(constant.CONFIG_PATH_ONE_SIGNAL_REST_API_KEY)

	if osAppId == "" {
		osAppId = constant.DEFAULT_ONE_SIGNAL_APP_ID
	}

	if osRestApiKey == "" {
		osRestApiKey = constant.DEFAULT_ONE_SIGNAL_REST_API_KEY
	}

	config := &dto.OneSignalConfig{
		AppId:      osAppId,
		RestApiKey: osRestApiKey,
	}

	log.Info("GetOneSignalConfig() - finished.")
	return config
}

func GetNewRelicConfig() *dto.NewRelicConfig {
	log.Info("GetNewRelicConfig() - starting...")
	osAppName := viper.GetString(constant.CONFIG_PATH_NEWRELIC_APP_NAME)
	osLicense := viper.GetString(constant.CONFIG_PATH_NEWRELIC_LICENSE)

	if osAppName == "" {
		osAppName = constant.DEFAULT_NEW_RELIC_APP_NAME
	}

	if osLicense == "" {
		osLicense = constant.DEFAULT_NEW_RELIC_LICENSE
	}

	config := &dto.NewRelicConfig{
		AppName: osAppName,
		License: osLicense,
	}

	log.Info("GetOneSignalConfig() - finished.")
	return config
}

func GetMinioConfig() *dto.MinioConfig {
	log.Info("GetMinioConfig() - starting...")
	minioHost := viper.GetString(constant.CONFIG_PATH_MINIO_HOST)
	minioLocation := viper.GetString(constant.CONFIG_PATH_MINIO_LOCATION)
	minioAccessKey := viper.GetString(constant.CONFIG_PATH_MINIO_ACCESS_KEY)
	minioSecretKey := viper.GetString(constant.CONFIG_PATH_MINIO_SECRET_KEY)
	minioSSL := viper.GetBool(constant.CONFIG_PATH_MINIO_SSL)
	minioReplaceDomain := viper.GetString(constant.CONFIG_PATH_MINIO_REPLACE_DOMAIN)

	if minioHost == "" {
		minioHost = constant.DEFAULT_MINIO_HOST
	}

	if minioLocation == "" {
		minioLocation = constant.DEFAULT_MINIO_LOCATION
	}

	if minioAccessKey == "" {
		minioAccessKey = constant.DEFAULT_MINIO_ACCESS_KEY
	}

	if minioSecretKey == "" {
		minioSecretKey = constant.DEFAULT_MINIO_SECRET_KEY
	}

	if minioReplaceDomain == "" {
		minioReplaceDomain = constant.DEFAULT_MINIO_REPLACE_DOMAIN
	}

	config := &dto.MinioConfig{
		Host:          minioHost,
		Location:      minioLocation,
		AccessKey:     minioAccessKey,
		SecretKey:     minioSecretKey,
		SSL:           minioSSL,
		ReplaceDomain: minioReplaceDomain,
	}

	log.Info("GetMinioConfig() - finished.")
	return config
}
