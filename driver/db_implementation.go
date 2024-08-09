package driver

import (
	"context"
	"fmt"
	"mf_backup_onetime/constant"
	"mf_backup_onetime/dto"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Implementation struct {
	Config *dto.DatabaseConfig
}

func InitImplementation() *Implementation {
	return &Implementation{}
}

func (i *Implementation) SetConfig(cfg *dto.DatabaseConfig) {
	i.Config = cfg
}

func (i *Implementation) NewMySql() (*gorm.DB, error) {
	var path string
	if i.Config.Dialect != "" {
		log.Infof("connect mysql : %v:%v", i.Config.Host, i.Config.Port)
		timeZone := strings.Replace(i.Config.TimeZone, "/", "%2F", -1)
		path = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=%v&charset=utf8&autocommit=false",
			i.Config.Username, i.Config.Password, i.Config.Host, i.Config.Port, i.Config.Database, timeZone)
	} else {
		log.Infof("connect mysql : %v:%v", constant.DEFAULT_DB_HOST, constant.DEFAULT_DB_PORT)
		timeZone := strings.Replace(constant.DEFAULT_TIME_ZONE, "/", "%2F", -1)
		path = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=%v&charset=utf8&autocommit=false",
			constant.DEFAULT_DB_USERNAME, constant.DEFAULT_DB_PASSWORD, constant.DEFAULT_DB_HOST,
			constant.DEFAULT_DB_PORT, constant.DEFAULT_DB_DATABASE, timeZone)
	}

	db, err := gorm.Open(mysql.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (i *Implementation) NewPostgreSql() (*gorm.DB, error) {
	var path string
	if i.Config.Dialect != "" {
		log.Infof("connect postgresql : %v:%v", i.Config.Host, i.Config.Port)
		path = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
			i.Config.Host, i.Config.Username, i.Config.Password, i.Config.Database, i.Config.Port, i.Config.SSLMode, i.Config.TimeZone)
	} else {
		log.Infof("connect postgresql : %v:%v", constant.DEFAULT_DB_HOST, constant.DEFAULT_DB_PORT)
		path = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
			constant.DEFAULT_DB_HOST, constant.DEFAULT_DB_USERNAME, constant.DEFAULT_DB_PASSWORD, constant.DEFAULT_DB_DATABASE,
			constant.DEFAULT_DB_PORT, constant.DEFAULT_DB_SSL_MODE, constant.DEFAULT_TIME_ZONE)
	}

	db, err := gorm.Open(postgres.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (i *Implementation) NewSqlServer() (*gorm.DB, error) {
	var path string
	if i.Config.Dialect != "" {
		log.Infof("connect sql server : %v:%v", i.Config.Host, i.Config.Port)
		path = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
			i.Config.Host, i.Config.Username, i.Config.Password, i.Config.Database, i.Config.Port, i.Config.SSLMode, i.Config.TimeZone)
	} else {
		log.Infof("connect sql server : %v:%v", constant.DEFAULT_DB_HOST, constant.DEFAULT_DB_PORT)
		path = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
			constant.DEFAULT_DB_HOST, constant.DEFAULT_DB_USERNAME, constant.DEFAULT_DB_PASSWORD, constant.DEFAULT_DB_DATABASE,
			constant.DEFAULT_DB_PORT, constant.DEFAULT_DB_SSL_MODE, constant.DEFAULT_TIME_ZONE)
	}

	db, err := gorm.Open(sqlserver.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (i *Implementation) NewSqlLite() (*gorm.DB, error) {
	var path string
	if i.Config.Dialect != "" {
		log.Infof("connect sqlite : %v:%v", i.Config.Host, i.Config.Port)
		path = fmt.Sprintf("%v", i.Config.Host)
	} else {
		log.Infof("connect sqlite : %v:%v", constant.DEFAULT_DB_HOST, constant.DEFAULT_DB_PORT)
		path = fmt.Sprintf("%v", constant.DEFAULT_PATH_SQLITE)
	}

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
func (i *Implementation) NewMongo(ctx context.Context) (*mongo.Database, error) {

	username, err := url.QueryUnescape(i.Config.Username)
	if err != nil {
		username = i.Config.Username
	}
	password, err := url.QueryUnescape(i.Config.Password)
	if err != nil {
		password = i.Config.Password
	}
	driver := "mongodb"
	URL := url.URL{
		Scheme: driver,
		Host:   fmt.Sprintf("%s:%s", i.Config.Host, i.Config.Port),

		Path:     i.Config.Database,
		User:     url.UserPassword(username, password),
		RawQuery: fmt.Sprintf("authSource=%v", i.Config.AuthSource),
	}

	fmt.Printf("URI ayam: %s \n", URL.String())
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

	return db.Database(i.Config.Database), nil
}
