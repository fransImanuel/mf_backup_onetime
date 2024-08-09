package driver

import (
	"context"
	"mf_backup_onetime/dto"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBMongo struct {
	DB                  *mongo.Database
	DatabaseConfigMongo *dto.DatabaseConfig
}

func InitDBMongo() *DBMongo {
	log.Info("InitDBMongo() - starting...")
	log.Info("InitDBMongo() - finished...")
	return &DBMongo{}
}

func (c *DBMongo) Factory(cfg *dto.DatabaseConfig, ctx context.Context) error {
	log.Info("DBMongo - Factory() - starting...")
	implementation := InitImplementation()
	implementation.SetConfig(cfg)

	dbConn, err := implementation.NewMongo(ctx)
	if err != nil {
		return err
	}
	c.DB = dbConn
	c.DatabaseConfigMongo = cfg

	log.Info("DBMongo - Factory() - finished.")
	return nil
}

func (c *DBMongo) Ping() error {
	log.Info("DBMongo - Ping() - starting...")
	if err := c.DB.Client().Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	log.Info("DBMongo - Ping() - finished.")
	return nil
}

func (c *DBMongo) Close() {
	log.Info("DBMongo - Close() - starting...")
	db := c.DB.Client()

	err := db.Disconnect(context.TODO())
	if err != nil {
		log.Fatal("DBMongo - Close() - error: ", err)
	}

	log.Info("DBMongo - Close() - finished.")
}
