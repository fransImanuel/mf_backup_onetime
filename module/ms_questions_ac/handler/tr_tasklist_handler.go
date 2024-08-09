package handler

import (
	"mf_backup_onetime/module/ms_questions_ac"
	"mf_backup_onetime/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MSQuestionHandler struct {
	MSQuestionService ms_questions_ac.Service
}

func InitMSQuestionHandler(tr_tasklist ms_questions_ac.Service) {
	// handler := &MSQuestionHandler{
	// 	MSQuestionService: tr_tasklist,
	// }

	// routeAPI := g.Group("/api/v1/question")

	// routeAPI.POST("/list", cHttp.Auth, handler.ListQuestionHandler)
	// routeAPI.GET("/mayapada", handler.QuestionMayapadaHandler)

}

// ListQuestionHandler
// @Tags MS Question
// @Summary ListQuestionHandler
// @Description ListQuestionHandler
// @ID ListQuestionHandler
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json

// @Success 200
// @Router /api/v1/question/list [post]
func (h *MSQuestionHandler) ListQuestionHandler(c *gin.Context) {

	List, err := h.MSQuestionService.ListQuestion(c)

	if err == nil {
		util.APIResponse(c, "List  successfully", http.StatusOK, nil, List)
	} else {

		util.APIResponse(c, "Internal Server Error", 500, nil, map[string][]interface{}{})
	}

}

// QuestionMayapadaHandler
// @Tags MS Question
// @Summary QuestionMayapadaHandler
// @Description QuestionMayapadaHandler
// @ID QuestionMayapadaHandler
// @Security ApiKeyAuth
// @Success 200
// @Router /v1/question/mayapada [get]
func (h *MSQuestionHandler) QuestionMayapadaHandler(c *gin.Context) {

	List, err := h.MSQuestionService.QuestionMayapada(c)

	if err == nil {
		util.APIResponse(c, "List successfully", http.StatusOK, nil, List)
	} else {
		//log.Errorln("error handler", err.Error())
		util.APIResponse(c, "Internal Server Error", 500, nil, map[string][]interface{}{})
	}

}
