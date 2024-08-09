package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DataPdf struct {
	Success    bool   `json:"success"`
	StatusCode string `json:"statusCode"`
	Message    string `json:"message"`
	Data       struct {
		TasklistId         string `json:"TasklistId"`
		DestinationCode    string `json:"DestinationCode"`
		DestinationName    string `json:"DestinationName"`
		DestinationAddress string `json:"DestinationAddress"`
		VendorName         string `json:"VendorName"`
		AssignedUser       string `json:"AssignedUser"`
		SurveyTime         string `json:"SurveyTime"`
		Status             string `json:"Status"`
		Duration           string `json:"Duration"`
		QuestionAnswer     []struct {
			Question struct {
				Id        string `json:"_id"`
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
			Answer struct {
				Id         string `json:"_id"`
				QuestionId string `json:"QuestionId"`
				Answer     struct {
					Value          *string `json:"Value"`
					ResultProperty struct {
						TakePhoto *string `json:"TakePhoto"`
					} `json:"ResultProperty"`
					ResultItem []struct {
						Value *string `json:"Value"`
					} `json:"ResultItem"`
				} `json:"Answer"`
			} `json:"Answer"`
		} `json:"QuestionAnswer"`
	} `json:"data"`
}

type Answer struct {
	Value          *string        `json:"Value"`
	ResultProperty ResultProperty `json:"ResultProperty"`
	ResultItem     []ResultItem   `json:"ResultItem"`
}

type ResultProperty struct {
	TakePhoto *string `json:"TakePhoto"`
}

type ResultItem struct {
	//Value *string `json:"Value"`
	Value interface{} `json:"value"`
}
type ResultSurvey struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id"`
	QuestionId primitive.ObjectID `bson:"QuestionId" json:"QuestionId"`
	Answer     Answer             `json:"Answer"`
}
type TasklistDetail struct {
	Id                 primitive.ObjectID `bson:"_id, omitempty"`
	SurveyTime         time.Time          `bson:"SurveyTime" json:"SurveyTime"`
	QuestionSurveyId   primitive.ObjectID `bson:"QuestionSurveyId" json:"QuestionSurveyId"`
	StatusId           int                `json:"StatusId"`
	DurationTime       string             `json:"DurationTime"`
	Latitude           string             `json:"Latitude"`
	Longitude          string             `json:"Longitude"`
	ResultSurvey       []ResultSurvey     `json:"ResultSurvey"`
	CreatedTime        time.Time          `bson:"CreatedTime, omitempty" json:"CreatedTime"`
	CreatedByUserName  string             `json:"CreatedByUserName"`
	ModifiedTime       time.Time          `bson:"ModifiedTime, omitempty" json:"ModifiedTime"`
	ModifiedByUserName string             `json:"ModifiedByUserName"`
}

type TRTasklist struct {
	Id                 primitive.ObjectID `bson:"_id, omitempty" json:"_id"`
	TenantId           int                `json:"TenantId"`
	DeletedByUserId    int                `json:"DeletedByUserId"`
	DeletedTime        time.Time          `bson:"DeletedTime, omitempty" `
	VisitId            int64              `json:"VisitId"`
	DestinationId      int64              `json:"DestinationId"`
	AssignedUserId     int64              `json:"AssignedUserId"`
	QuestionSurveyId   primitive.ObjectID `bson:"QuestionSurveyId" json:"QuestionSurveyId"`
	ScheduleVisit      time.Time          `bson:"ScheduleVisit, omitempty" json:"ScheduleVisit"`
	StatusId           int                `json:"StatusId"`
	TasklistDetail     []*TasklistDetail  `json:"TasklistDetail"`
	IsActive           bool               `json:"IsActive"`
	IsDeleted          bool               `json:"IsDeleted"`
	CreatedByUserName  string             `json:"CreatedByUserName"`
	CreatedTime        time.Time          `bson:"CreatedTime, omitempty" json:"CreatedTime"`
	ModifiedByUserName string             `bson:"ModifiedByUserName, omitempty" json:"ModifiedByUserName"`
	ModifiedTime       time.Time          `bson:"ModifiedTime, omitempty" json:"ModifiedTime"`
	V                  int                `json:"__v"`
}

type MongoTRTasklistAC struct {
	Id                 primitive.ObjectID `bson:"_id, omitempty"`
	TenantId           int                `bson:"TenantId, omitempty"`
	DeletedByUserId    int                `bson:"DeletedByUserId, omitempty"`
	DeletedTime        time.Time          `bson:"DeletedTime, omitempty" `
	VisitId            int                `bson:"VisitId, omitempty"`
	DestinationId      int64              `bson:"DestinationId, omitempty"`
	AssignedUserId     int64              `bson:"AssignedUserId, omitempty"`
	QuestionSurveyId   primitive.ObjectID `bson:"QuestionSurveyId, omitempty"`
	ScheduleVisit      time.Time          `bson:"ScheduleVisit, omitempty"`
	StatusId           int                `bson:"StatusId, omitempty"`
	TasklistDetail     []interface{}      `bson:"TasklistDetail, omitempty"`
	IsActive           *bool              `bson:"IsActive, omitempty"`
	IsDeleted          *bool              `bson:"IsDeleted, omitempty"`
	CreatedByUserName  string             `bson:"CreatedByUserName, omitempty"`
	CreatedTime        time.Time          `bson:"CreatedTime, omitempty"`
	ModifiedByUserName string             `bson:"ModifiedByUserName, omitempty"`
	ModifiedTime       time.Time          `bson:"ModifiedTime, omitempty"`
	V                  int                `json:"__v"`
}

type ReportTRTasklistAC struct {
	Id    int64 `bson:"_id, omitempty"`
	Count int64 `bson:"count, omitempty"`
	//Data []MongoTRTasklistAC `bson:"data, omitempty"`
	DestinationId []int64             `bson:"DestinationId, omitempty"`
	DestId        []int64             `bson:"dest_id, omitempty"`
	DestIdLength  int64               `bson:"dest_id_length, omitempty"`
	Data          []DataSummaryStatus `bson:"data, omitempty"`
	//Date []time.Time `bson:"Date, omitempty"`
}
type DataSummaryStatus struct {
	TasklistId primitive.ObjectID `bson:"tasklist_id, omitempty"`
	StatusId   int                `bson:"status_id, omitempty"`
}
type CountReportTRTasklistAC struct {
	Id    int64 `bson:"_id, omitempty"`
	Count int64 `bson:"count, omitempty"`
}

type SchemaReportGenerate struct {
	Id     int64        `bson:"_id, omitempty"`
	Count  int64        `bson:"count, omitempty"`
	DestId []int64      `bson:"dest_id, omitempty"`
	Items  []TRTasklist `bson:"items, omitempty"`
}
type ListVisitATM struct {
	Id             primitive.ObjectID `bson:"_id, omitempty" json:"_id"`
	VisitId        int64              `json:"VisitId"`
	AssignedUserId int64              `json:"AssignedUserId"`
	StatusId       int                `json:"StatusId"`
	ScheduleVisit  time.Time          `bson:"ScheduleVisit, omitempty" json:"ScheduleVisit"`
}

type ReportGenerateResponse struct {
	CountVisit int        `bson:"CountVisit"`
	VisitId    int64      `bson:"VisitId"`
	AssignTime *time.Time `json:"assign_time" `
	StartVisit *time.Time `json:"start_visit" `

	DestinationId      int64  `json:"DestinationId" `
	DestinationCode    string `json:"DestinationCode" `
	DestinationName    string `json:"DestinationName" `
	DestinationAddress string `json:"DestinationAddress" `

	VendorId   int64  `json:"VendorId" `
	VendorName string `json:"VendorName" `

	AssignedUserId   int64  `json:"AssignedUserId" `
	AssignedUserName string `json:"AssignedUserName"`

	Tasklist []ListVisitATM `bson:"Tasklist, omitempty"`
}
