package dto

import (
	"mf_backup_onetime/module/tr_tasklist/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TRTaskListRequest struct {
	FilterBaseDto
	UserId        []int64   `json:"user_id"`
	IsDeleted     *bool     `json:"IsDeleted"`
	IsActive      *bool     `json:"IsActive"`
	ScheduleVisit time.Time `json:"ScheduleVisit"`
	StatusId      []int     `json:"StatusId"`
	StartDate     string    `json:"start_date"`
	EndDate       string    `json:"end_date"`
	DestinationId []int64   `json:"destination_id"`

	//model.MongoTRTasklistAC
}
type TRTaskListRequest2 struct {
	FilterBaseDto
	UserId        []int64   `json:"user_id"`
	IsDeleted     *bool     `json:"IsDeleted"`
	IsActive      *bool     `json:"IsActive"`
	ScheduleVisit time.Time `json:"ScheduleVisit"`
	StatusId      []int     `json:"StatusId"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	DestinationId int64     `json:"destination_id"`
	//model.MongoTRTasklistAC
}

type ReportTaskListRequest struct {
	VendorName    string `json:"vendor_name"`
	JumlahATM     int64  `json:"jumlah_atm"`
	Date          string `json:"date"`
	Name          string `json:"name"`
	TotalTasklist int64  `json:"total_tasklist"`
	StatusDone    int64  `json:"status_done"`
	StatusDelay   int64  `json:"status_delay"`
	StatusOpen    int64  `json:"status_open"`
	StatusManual  int64  `json:"status_manual"`
	//DestinationId string `json:"destination_id"`
	//DestinationName string `json:"destination_name"`

}

type VendorAndATMResponse struct {
	TasklistId      primitive.ObjectID `bson:"TasklistId,omitempty" json:"TasklistId,omitempty"`
	UserId          int64              `json:"UserId,omitempty"`
	FullName        string             `json:"FullName,omitempty"`
	VendorId        int64              `json:"VendorId,omitempty"`
	VendorName      string             `json:"VendorName,omitempty"`
	StatusId        int                `json:"StatusId,omitempty"`
	StatusName      string             `json:"StatusName,omitempty"`
	ScheduleVisit   time.Time          `json:"ScheduleVisit,omitempty"`
	SurveyTime      string             `json:"SurveyTime,omitempty"`
	DurationTime    string             `json:"DurationTime,omitempty"`
	DestinationId   int64              `json:"DestinationId,omitempty"`
	Destination     string             `json:"Destination,omitempty"`
	DestinationCode string             `json:"DestinationCode,omitempty"`
}

type ReportPDFList struct {
	TasklistId         primitive.ObjectID `bson:"TasklistId"`
	DestinationCode    string             `json:"DestinationCode"`
	DestinationName    string             `json:"DestinationName"`
	DestinationAddress string             `json:"DestinationAddress"`
	VendorName         string             `json:"VendorName"`
	AssignedUser       string             `json:"AssignedUser"`
	ScheduleVisit      time.Time          `json:"ScheduleVisit"`
	SurveyTime         time.Time          `json:"SurveyTime"`
	Status             string             `json:"Status"`
	Duration           string             `json:"Duration"`
	QuestionAnswer     []*QuestionAnswer  `json:"QuestionAnswer"`
}

type QuestionAnswer struct {
	Question Question           `json:"Question"`
	Answer   model.ResultSurvey `json:"Answer"`
}

type Question struct {
	Id        primitive.ObjectID `bson:"_id, omitempty"`
	TypeField string             `json:"TypeField"`
	Label     string             `json:"Label"`
	Item      []Item             `json:"Item"`
}
type Item struct {
	Id        int     `json:"id"`
	Key       string  `json:"key"`
	Value     int     `json:"value"`
	Action    *string `json:"action"`
	Condition string  `json:"condition"`
}
type ReportRequest struct {
	FilterBaseDto
	VendorId []int64 `json:"vendor_id"`

	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type GetTasklistByID struct {
	TasklistId string `json:"tasklist_id"`
}

type VendorAndATMRequest struct {
	FilterBaseDto
	DestinationId int64   `json:"destination_id"`
	UserId        []int64 `json:"user_id"`
	//StartDate string `json:"start_date" example:"2022-09-01"`
	Period string `json:"period" example:"2022-09"`
	//EndDate   string `json:"end_date" example:"2022-09-30"`
}

type ReportResponseDto struct {
	ResponseBaseDto
	Id    int                     `json:"id,omitempty"`
	Count int64                   `json:"count,omitempty"`
	Items []ReportTaskListRequest `json:"items,omitempty"`

	//Item  ReportTaskListRequest   `json:"item,omitempty"`
}
type ReportPDFResponseDto struct {
	ResponseBaseDto
	Id    int              `json:"id,omitempty"`
	Count int64            `json:"count,omitempty"`
	Items []*ReportPDFList `json:"items,omitempty"`

	//Item  ReportTaskListRequest   `json:"item,omitempty"`
}

type NegativeCondition struct {
	QuestionId primitive.ObjectID `bson:"QuestionId"`
	Value      *string            `bson:"Answer.ResultItem.Value"`
}

type ParamNegativeCondition struct {
	QuestionId string `json:"QuestionId" example:"60f1b1b9b1b1b1b1b1b1b1b1"`
	Value      int    `json:"Value" example:"1"`
}

type SchemaNegativeCondition struct {
	StartDate         string                   `json:"startDate" example:"2022-09-01"`
	EndDate           string                   `json:"endDate" example:"2022-09-01"`
	VendorId          int                      `json:"vendorId" example:"1"`
	VendorName        string                   `json:"vendorName" example:"Labil"`
	NegativeCondition []ParamNegativeCondition `json:"negativeCondition"`
	//IsExport interface{} `json:"isExport"`
	Offset int64 `json:"offset" example:"0"`
	Limit  int64 `json:"limit" example:"10"`
}

type TasklistNegativeCondition struct {
	Id             primitive.ObjectID `bson:"_id, omitempty"`
	DestinationId  int64              `json:"DestinationId"`
	AssignedUserId int64              `json:"AssignedUserId"`
	ScheduleVisit  time.Time          `json:"ScheduleVisit"`
	StatusId       int                `json:"StatusId"`
	TasklistDetail []struct {
		SurveyTime   time.Time `json:"SurveyTime"`
		DurationTime string    `json:"DurationTime"`
	} `json:"TasklistDetail"`
}

type TasklistNegativeConditionResponse struct {
	TasklistId         primitive.ObjectID `json:"TasklistId"`
	FullName           string             `json:"FullName"`
	DestinationCode    string             `json:"DestinationCode"`
	DestinationName    string             `json:"DestinationName"`
	DestinationAddress string             `json:"DestinationAddress"`
	ScheduleVisit      time.Time          `json:"ScheduleVisit"`
	StatusId           int                `json:"StatusId"`
	StatusName         string             `json:"StatusName"`
	SurveyTime         time.Time          `json:"SurveyTime"`
	DurationTime       string             `json:"DurationTime"`
	DurationFlag       bool               `json:"DurationFlag"`
}

type ResponseNegativeConditionDto struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	ItemsCount int    `json:"itemsCount"`
}

type SchemaCheckDataDoubleResponse struct {
	Id     int64              `bson:"_id, omitempty" json:"_id"`
	Count  int64              `bson:"count" json:"count"`
	DestId []int64            `json:"dest_id"`
	Data   []model.TRTasklist `json:"data"`
}
type SchemaCheckDataDoubleRequest struct {
	StartDate string `json:"startDate" binding:"required" example:"2022-09-01"`
	Offset    int64  `json:"offset" example:"0"`
	Limit     int64  `json:"limit" example:"10"`
}
type ReportGenerateRequest struct {
	FilterPage        int     `json:"filter_page" example:"1"`
	FilterLimit       int     `json:"filter_limit" example:"10"`
	VendorId          int64   `json:"VendorId" example:"1"`
	DestinationId     int64   `json:"DestinationId" example:"123"`
	Period            string  `json:"Period" example:"2022-09"`
	NotAssignedUserId []int64 `json:"NotAssignedUserId" example:"1"`
	Date              int     `json:"Date" example:"1"`
}
type ReportGenerateRequestRepository struct {
	UserId        []int64   `json:"UserId"`
	ScheduleVisit time.Time `json:"ScheduleVisit"`
	StatusId      []int     `json:"StatusId"`
	StartDate     time.Time `json:"StartDate"`
	EndDate       time.Time `json:"EndDate"`
	VisitId       []int64   `json:"VisitId"`
	FilterPage    int       `json:"filter_page" example:"1"`
	FilterLimit   int       `json:"filter_limit" example:"10"`
}
