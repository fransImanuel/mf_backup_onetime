package repository

import (
	"context"
	"mf_backup_onetime/constant"
	"mf_backup_onetime/module/ms_questions_ac"
	"mf_backup_onetime/module/ms_questions_ac/model"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MSQuestionRepository struct {
	DBMongo *mongo.Database
}

func InitMSQuestionRepository(db *mongo.Database) ms_questions_ac.Repository {
	return &MSQuestionRepository{
		DBMongo: db,
	}
}

func (r *MSQuestionRepository) Get(ctx context.Context) (res []*model.MSQuestionSurveysAC, count int64, err error) {
	log.Info("MSQuestionRepository - Get() - starting...")

	db := r.DBMongo

	//r.DBMongo.Client()
	if err != nil {
		log.Fatal(err.Error())
	}

	findOptions := options.Find()

	//orderByQuery := 1
	//if req.OrderField != "" {
	//	orderColumn, orderType := util.SplitOrderQuery(req.OrderField)
	//
	//	if orderType == "asc" {
	//		orderByQuery = -1
	//	} else {
	//		orderByQuery = 1
	//	}
	//
	//	switch orderColumn {
	//	case "updated_at":
	//		findOptions.SetSort(bson.D{{"updated_at", orderByQuery}})
	//	case "date":
	//		findOptions.SetSort(bson.D{{"date", orderByQuery}})
	//	default:
	//		findOptions.SetSort(bson.D{{"created_at", orderByQuery}})
	//	}
	//} else {
	//	findOptions.SetSort(bson.D{{"created_at", -1}})
	//}

	//if req.FilterPage > 0 && req.FilterLimit > 0 {
	//	offset := int64(req.FilterPage*req.FilterLimit - req.FilterLimit)
	//	limit := int64(req.FilterLimit)
	//	findOptions.Limit = &limit
	//	findOptions.Skip = &offset
	//}

	csr, err := db.Collection(constant.MONGO_COLLECTION_MS_QUESTION_SURVEYS_AC).Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]*model.MSQuestionSurveysAC, 0)
	for csr.Next(ctx) {
		var row *model.MSQuestionSurveysAC
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	res = result

	log.Info("MSQuestionRepository - Get() - finished.")
	return res, count, nil
}

func (r *MSQuestionRepository) GetMSQuestionSurveys(ctx context.Context) (res *model.MSQuestionSurveys, err error) {
	log.Info("MSQuestionRepository - GetMSQuestionSurveys() - starting...")

	db := r.DBMongo

	var MSQuestionSurveys *model.MSQuestionSurveys
	err = db.Collection(constant.MONGO_COLLECTION_MS_QUESTION_SURVEYS_AC).FindOne(ctx, bson.M{}).Decode(&MSQuestionSurveys)
	if err != nil {
		// log.Fatal(err.Error())
		return nil, err
	}

	res = MSQuestionSurveys

	log.Info("MSQuestionRepository - GetMSQuestionSurveys() - finished.")
	return res, nil
}
func (r *MSQuestionRepository) GetMSQuestionSurveysMayapada(ctx context.Context) (res *model.MSQuestionSurveys, err error) {
	log.Info("MSQuestionRepository - GetMSQuestionSurveysMayapada() - starting...")

	db := r.DBMongo

	var MSQuestionSurveys *model.MSQuestionSurveys
	err = db.Collection(constant.MONGO_COLLECTION_MS_QUESTION_SURVEYS_MAYAPADA).FindOne(ctx, bson.M{}).Decode(&MSQuestionSurveys)
	if err != nil {
		log.Println("err.Error()", err.Error())
	}

	res = MSQuestionSurveys

	log.Info("MSQuestionRepository - GetMSQuestionSurveysMayapada() - finished.")
	return res, nil
}
