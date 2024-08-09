package model

import (
	"mf_backup_onetime/dto"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MSQuestionSurveysAC struct {
	Id                 primitive.ObjectID `bson:"_id, omitempty"`
	Code               string             `json:"Code"`
	CreatedByUserName  string             `json:"CreatedByUserName"`
	CreatedTime        time.Time          `bson:"CreatedTime"`
	DeletedByUserId    int                `json:"DeletedByUserId"`
	DeletedTime        time.Time          `bson:"DeletedTime" json:"DeletedTime"`
	Information        string             `json:"Information"`
	IsActive           bool               `json:"IsActive"`
	IsDeleted          bool               `json:"IsDeleted"`
	ModifiedByUserName string             `json:"ModifiedByUserName"`
	ModifiedTime       time.Time          `bson:"ModifiedTime"`
	Name               string             `json:"Name"`
	Question           []struct {
		Id struct {
			Oid string `json:"$oid"`
		} `json:"_id"`
		TypeField string `json:"TypeField"`
		Label     string `json:"Label"`
		Item      []struct {
			Id        int     `json:"id"`
			Key       string  `json:"key"`
			Value     int     `json:"value"`
			Action    *string `json:"action"`
			Condition string  `json:"condition"`
		} `json:"Item"`
	} `json:"Question"`
	TenantId int `json:"TenantId"`
}

type MSQuestionSurveys struct {
	Id                 primitive.ObjectID `bson:"_id, omitempty"`
	Code               string             `json:"Code"`
	CreatedByUserName  string             `json:"CreatedByUserName"`
	CreatedTime        time.Time          `bson:"CreatedTime"`
	DeletedByUserId    int                `json:"DeletedByUserId"`
	DeletedTime        time.Time          `bson:"DeletedTime" json:"DeletedTime"`
	Information        string             `json:"Information"`
	IsActive           bool               `json:"IsActive"`
	IsDeleted          bool               `json:"IsDeleted"`
	ModifiedByUserName string             `json:"ModifiedByUserName"`
	ModifiedTime       time.Time          `bson:"ModifiedTime"`
	Name               string             `json:"Name"`
	Question           []dto.Question     `json:"Question"`
	TenantId           int                `json:"TenantId"`
}

//type Question struct {
//	Id   primitive.ObjectID `bson:"_id, omitempty"`
//	TypeField string `json:"TypeField"`
//	Label     string `json:"Label"`
//	Item      []Item `json:"Item"`
//}
//type Item struct {
//	Id        int     `json:"id"`
//	Key       string  `json:"key"`
//	Value     int     `json:"value"`
//	Action    *string `json:"action"`
//	Condition string  `json:"condition"`
//}
