package ms_questions_ac

import (
	"context"
	"mf_backup_onetime/module/ms_questions_ac/model"
)

type Repository interface {
	Get(ctx context.Context) (res []*model.MSQuestionSurveysAC, count int64, err error)
	GetMSQuestionSurveys(ctx context.Context) (res *model.MSQuestionSurveys, err error)
	GetMSQuestionSurveysMayapada(ctx context.Context) (res *model.MSQuestionSurveys, err error)
	//Insert(ctx context.Context,  req model.MongoTRTasklistAC) (err error)
}

type Service interface {
	ListQuestion(ctx context.Context) (res []*model.MSQuestionSurveysAC, err error)
	QuestionMayapada(ctx context.Context) (res *model.MSQuestionSurveys, err error)
}
