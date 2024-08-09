package driver

import (
	"context"
	"fmt"
	"mf_backup_onetime/schemas"
	"net/url"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func SetupMongodb(ctx context.Context, config schemas.SchemaEnvironment) (*mongo.Database, error) {
	log.Info("GetDatabaseConfig() - starting...")
	dbHost := config.MONGO_HOST
	dbPort := config.MONGO_PORT
	dbUsername := config.MONGO_USER
	dbPassword := config.MONGO_PASS
	dbName := config.MONGO_DB
	//dbSSLMode := viper.GetString(constant.CONFIG_PATH_DB_SSL_MODE_MONGO)
	dbAuthSource := config.MONGO_AUTH
	//dbTimeZone := viper.GetString(constant.CONFIG_PATH_APP_TIME_ZONE)

	username, err := url.QueryUnescape(dbUsername)
	if err != nil {
		username = dbUsername
	}
	password, err := url.QueryUnescape(dbPassword)
	if err != nil {
		password = dbPassword
	}
	driver := "mongodb"
	URL := url.URL{
		Scheme: driver,
		Host:   fmt.Sprintf("%s:%s", dbHost, dbPort),

		Path:     dbName,
		User:     url.UserPassword(username, password),
		RawQuery: fmt.Sprintf("authSource=%v", dbAuthSource),
	}

	fmt.Printf("URI MongoDB: %s \n", URL.String())
	//assumeRoleCredential := options.Credential{
	//	Username: username,
	//	Password: password,
	//}
	clientOptions := options.Client()
	//clientOptions := options.Client().SetAuth(assumeRoleCredential)
	//clientOptions.ApplyURI(fmt.Sprintf("mongodb://%v:%v", i.Config.Host, i.Config.Port))
	clientOptions.ApplyURI(URL.String())
	clientOptions.SetDirect(true)
	db, err := mongo.NewClient(clientOptions)
	if err != nil {

		return nil, err
	}

	err = db.Connect(ctx)
	if err != nil {
		fmt.Printf("error when mongo.Connect(ctx, clientOptions), %v, %v", clientOptions, err)
		return nil, err
	}
	err = db.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	//defer db.Disconnect(context.Background())
	fmt.Println("💚 Connected to MongoDB!")

	return db.Database(dbName), nil

}
