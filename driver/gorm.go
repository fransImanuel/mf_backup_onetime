package driver

import (
	"fmt"
	"mf_backup_onetime/constant"
	"mf_backup_onetime/dto"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DbContext struct {
	DB             *gorm.DB
	DatabaseConfig *dto.DatabaseConfig
}

func InitDbContext() *DbContext {
	log.Info("InitDbContext() - starting...")
	log.Info("InitDbContext() - finished...")
	return &DbContext{}
}

func (c *DbContext) Factory(cfg *dto.DatabaseConfig) error {
	log.Info("DbContext - Factory() - starting...")
	implementation := InitImplementation()
	implementation.SetConfig(cfg)

	switch cfg.Dialect {
	case constant.DIALECT_MYSQL:
		dbConn, err := implementation.NewMySql()
		if err != nil {
			return err
		}
		c.DB = dbConn
	case constant.DIALECT_POSTGRESQL:
		dbConn, err := implementation.NewPostgreSql()
		if err != nil {
			return err
		}
		c.DB = dbConn
	case constant.DIALECT_SQL_SERVER:
		dbConn, err := implementation.NewSqlServer()
		if err != nil {
			return err
		}
		c.DB = dbConn
	case constant.DIALECT_SQLITE:
		dbConn, err := implementation.NewSqlLite()
		if err != nil {
			return err
		}
		c.DB = dbConn
	default:
		return fmt.Errorf("dialect not found")
	}

	c.DatabaseConfig = cfg
	log.Info("DbContext - Factory() - finished.")
	return nil
}

func (c *DbContext) Ping() error {
	log.Info("DbContext - Ping() - starting...")
	db, err := c.DB.DB()
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	log.Info("DbContext - Ping() - finished.")
	return nil
}

func (c *DbContext) Close() {
	log.Info("DbContext - Close() - starting...")
	db, err := c.DB.DB()
	if err != nil {
		log.Fatal("DbContext - Close() - error: ", err)
	}

	err = db.Close()
	if err != nil {
		log.Fatal("DbContext - Close() - error: ", err)
	}

	log.Info("DbContext - Close() - finished.")
}
