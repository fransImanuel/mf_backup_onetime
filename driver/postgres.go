package driver

import (
	"fmt"
	"mf_backup_onetime/schemas"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDBSQL(config schemas.SchemaEnvironment) (*gorm.DB, error) {
	logrus.Debug("🔌 Connecting into Database Postgres")
	dbHost := config.DB_HOST
	dbUsername := config.DB_USER
	dbPassword := config.DB_PASS
	dbName := config.DB_NAME
	dbPort := config.DB_PORT
	dbSSLMode := config.DB_SSLMODE
	timezone := config.TIMEZONE

	path := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		dbHost, dbUsername, dbPassword, dbName, dbPort, dbSSLMode, timezone)

	db, err := gorm.Open(postgres.Open(path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		defer logrus.Errorln("❌ Error Connect into Database Postgres", err.Error())

		return nil, err
	}

	postgreSQL, err := db.DB()
	// Set connection pool parameters
	postgreSQL.SetMaxOpenConns(10)   // Maximum number of open connections
	postgreSQL.SetMaxIdleConns(5)    // Maximum number of idle connections
	postgreSQL.SetConnMaxLifetime(0) // Connection lifetime (0 means connections are reused indefinitely)

	if os.Getenv("GO_ENV") == "development" {
		db.Debug()
	}
	fmt.Println("💚 Connect into Database Postgres Success")

	return db, nil
}
