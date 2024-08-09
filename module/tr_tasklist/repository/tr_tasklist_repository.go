package repository

import (
	"context"
	"fmt"
	"mf_backup_onetime/module/tr_tasklist"
	"mf_backup_onetime/module/tr_tasklist/model"
	"time"

	"mf_backup_onetime/constant"
	"mf_backup_onetime/schemas"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type TRTasklistRepository struct {
	DBMongo    *mongo.Database
	DBPostgres *gorm.DB
}

func InitTRTasklistRepository(dbConMonggo *mongo.Database, DBPostgres *gorm.DB) tr_tasklist.Repository {
	return &TRTasklistRepository{
		DBMongo:    dbConMonggo,
		DBPostgres: DBPostgres,
	}
}

func (r *TRTasklistRepository) BulkMongoExportOneTime() []string {
	log.Debug("TRTasklistRepository - UpdateScheduleVisitFixGenerateRepo() - starting...")
	// Define the date range
	startOfMay := time.Date(2024, time.May, 1, 0, 0, 0, 0, time.UTC)
	endOfMay := time.Date(2024, time.June, 1, 0, 0, 0, 0, time.UTC)
	// startOfJune := time.Date(2024, time.June, 1, 0, 0, 0, 0, time.UTC)
	// endOfJune := time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC)
	// SecondOfMay := time.Date(2024, time.May, 2, 0, 0, 0, 0, time.UTC)

	// Create the filter
	filter := bson.M{
		"ScheduleVisit": bson.M{
			"$gte": startOfMay,
			"$lt":  endOfMay,
		},
	}

	// filter := bson.M{}

	db := r.DBMongo

	// Select the database and collection
	collection := db.Collection("tr_tasklists")

	// Execute the query
	// cursor, err := collection.Find(context.TODO(), filter)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return nil
	}
	defer cursor.Close(context.TODO())

	type TRTasklistID struct {
		ID string `json:"_id" bson:"_id"`
	}
	IDs := []string{}
	// Iterate over the cursor
	for cursor.Next(context.TODO()) {
		// var document bson.M
		var document TRTasklistID
		err := cursor.Decode(&document)
		if err != nil {
			fmt.Println("Failed to decode document:", err)
			return nil
		}
		// fmt.Println(document)
		IDs = append(IDs, document.ID)
	}

	if err := cursor.Err(); err != nil {
		fmt.Println("Cursor error:", err)
		return nil
	}

	log.Debug("TRTasklistRepository - UpdateScheduleVisitFixGenerateRepo() - finished...")
	return IDs
}

func (r *TRTasklistRepository) GetTasklistByIdRepository(id primitive.ObjectID) (model.TRTasklist, schemas.SchemaDatabaseError) {

	DBMongo := r.DBMongo

	var DataTasklist model.TRTasklist
	Find := DBMongo.Collection(constant.MONGO_COLLECTION_TR_TASKLISTS).FindOne(context.Background(), bson.M{"_id": id}).Decode(&DataTasklist)
	if Find != nil {
		log.Errorln(Find.Error())
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	DataTasklist.ScheduleVisit.In(loc)

	for _, item := range DataTasklist.TasklistDetail {
		item.SurveyTime.In(loc)
	}

	return DataTasklist, schemas.SchemaDatabaseError{}
}
