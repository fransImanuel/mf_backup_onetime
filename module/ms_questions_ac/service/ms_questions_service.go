package service

import (
	"context"
	"fmt"
	"mf_backup_onetime/module/ms_questions_ac"
	"mf_backup_onetime/module/ms_questions_ac/model"

	log "github.com/sirupsen/logrus"
)

type MSQuestionService struct {
	MSQuestionRepository ms_questions_ac.Repository
}

func InitMSQuestionService(ms_questions_ac ms_questions_ac.Repository) ms_questions_ac.Service {
	return &MSQuestionService{

		MSQuestionRepository: ms_questions_ac,
	}
}

func (s *MSQuestionService) ListQuestion(ctx context.Context) (res []*model.MSQuestionSurveysAC, err error) {
	log.Info("MSQuestionService - ListQuestion() - starting...")

	result, _, err := s.MSQuestionRepository.Get(ctx)

	res = result
	log.Info("MSQuestionService - ListQuestion() - finished.")
	return res, nil
}
func (s *MSQuestionService) QuestionMayapada(ctx context.Context) (res *model.MSQuestionSurveys, err error) {
	log.Info("MSQuestionService - QuestionMayapada() - starting...")

	result, err := s.MSQuestionRepository.GetMSQuestionSurveysMayapada(ctx)

	fmt.Println("err service", err)
	res = result
	log.Info("MSQuestionService - QuestionMayapada() - finished.")
	return res, nil
}
