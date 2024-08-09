package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"mf_backup_onetime/driver"
	"mf_backup_onetime/schemas"

	msDestinationRepository "mf_backup_onetime/module/ms_destination/repository"

	msUserRepository "mf_backup_onetime/module/ms_user/repository"

	ms_QuestionHandler "mf_backup_onetime/module/ms_questions_ac/handler"
	ms_QuestionRepository "mf_backup_onetime/module/ms_questions_ac/repository"
	ms_QuestionService "mf_backup_onetime/module/ms_questions_ac/service"

	trTasklistHandler "mf_backup_onetime/module/tr_tasklist/handler"
	trTasklistRepository "mf_backup_onetime/module/tr_tasklist/repository"
	trTasklistService "mf_backup_onetime/module/tr_tasklist/service"

	"github.com/joho/godotenv"
)

var ConfigEnv schemas.SchemaEnvironment

func main() {

	ConfigEnv = Environment()

	// client, err := connectToMongoDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	DBPostgres, err := driver.SetupDBSQL(ConfigEnv)
	if err != nil {
		log.Fatal(err.Error())
	}

	DBMongo, err := driver.SetupMongodb(context.Background(), ConfigEnv)
	if err != nil {
		log.Fatal(err.Error())
	}

	msUserRepo := msUserRepository.InitMSUserRepository(DBPostgres)
	msDestinationRepo := msDestinationRepository.InitMsDestinationRepository(DBPostgres)

	msQuestionRepo := ms_QuestionRepository.InitMSQuestionRepository(DBMongo)
	msQuestionServ := ms_QuestionService.InitMSQuestionService(msQuestionRepo)
	ms_QuestionHandler.InitMSQuestionHandler(msQuestionServ)

	trTasklistRepo := trTasklistRepository.InitTRTasklistRepository(DBMongo, DBPostgres)
	trTasklistServ := trTasklistService.InitTRTasklistService(trTasklistRepo, msQuestionRepo, msDestinationRepo, msUserRepo)
	trTasklistHandler.InitTRTasklistHandler(trTasklistServ)

}

// func connectToMongoDB() (*mongo.Client, error) {
// 	// Set client options
// 	clientOptions := options.Client().ApplyURI("mongodb://mongoAdmin:&Mer4h&Mud4&@10.254.213.3:4949")

// 	// Connect to MongoDB
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Check the connection
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("Connected to MongoDB!")
// 	return client, nil
// }

// func BulkMongoExportOneTime(collection *mongo.Collection) []string {
// 	fmt.Println("TRTasklistRepository - UpdateScheduleVisitFixGenerateRepo() - starting...")
// 	// Define the date range
// 	startOfMay := time.Date(2024, time.May, 1, 0, 0, 0, 0, time.UTC)
// 	SecondOfMay := time.Date(2024, time.May, 2, 0, 0, 0, 0, time.UTC)
// 	// endOfMay := time.Date(2024, time.June, 1, 0, 0, 0, 0, time.UTC)

// 	// Create the filter
// 	filter := bson.M{
// 		"ScheduleVisit": bson.M{
// 			"$gte": startOfMay,
// 			"$lt":  SecondOfMay,
// 		},
// 	}

// 	// filter := bson.M{}

// 	// Execute the query
// 	// cursor, err := collection.Find(context.TODO(), filter)
// 	cursor, err := collection.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Println("Failed to execute query:", err)
// 		return nil
// 	}
// 	defer cursor.Close(context.TODO())

// 	type TRTasklistID struct {
// 		ID string `json:"_id" bson:"_id"`
// 	}
// 	IDs := []string{}
// 	// Iterate over the cursor
// 	for cursor.Next(context.TODO()) {
// 		// var document bson.M
// 		var document TRTasklistID
// 		err := cursor.Decode(&document)
// 		if err != nil {
// 			fmt.Println("Failed to decode document:", err)
// 			return nil
// 		}
// 		// fmt.Println(document)
// 		IDs = append(IDs, document.ID)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		fmt.Println("Cursor error:", err)
// 		return nil
// 	}

// 	fmt.Println("TRTasklistRepository - UpdateScheduleVisitFixGenerateRepo() - finished...")
// 	return IDs
// }

func Environment() (config schemas.SchemaEnvironment) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Read environment variables from docker-compose.yml
	config.DB_HOST = os.Getenv("DB_HOST")
	config.DB_PORT = os.Getenv("DB_PORT")
	config.DB_USER = os.Getenv("DB_USER")
	config.DB_NAME = os.Getenv("DB_NAME")
	config.DB_PASS = os.Getenv("DB_PASS")
	config.DB_SSLMODE = os.Getenv("DB_SSLMODE")

	config.MONGO_HOST = os.Getenv("MONGO_HOST")
	config.MONGO_PORT = os.Getenv("MONGO_PORT")
	config.MONGO_USER = os.Getenv("MONGO_USER")
	config.MONGO_PASS = os.Getenv("MONGO_PASS")
	config.MONGO_DB = os.Getenv("MONGO_DB")
	config.MONGO_SSL = os.Getenv("MONGO_SSL")
	config.MONGO_AUTH = os.Getenv("MONGO_AUTH")

	config.TIMEZONE = os.Getenv("TIMEZONE")
	config.REST_PORT = os.Getenv("REST_PORT")
	config.GO_ENV = os.Getenv("GO_ENV")
	config.SWAGGER_HOST = os.Getenv("SWAGGER_HOST")
	config.JWT_SECRET = os.Getenv("JWT_SECRET")
	config.DOMAIN_IMAGE = os.Getenv("DOMAIN_IMAGE")

	config.Minio_Host = os.Getenv("MINIO_HOST")
	config.Minio_SSL = os.Getenv("MINIO_SSL")
	config.Minio_SecretKey = os.Getenv("MINIO_SECRET_KEY")
	config.Minio_AccessKey = os.Getenv("MINIO_ACCESS_KEY")
	config.Minio_Domain = os.Getenv("MINIO_DOMAIN")

	fmt.Printf("Environment sini: %+v\n", config)

	return config
}
