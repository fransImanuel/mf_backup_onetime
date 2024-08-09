package handler

import (
	"fmt"
	"mf_backup_onetime/dto"
	cHttp "mf_backup_onetime/http"
	"mf_backup_onetime/module/ms_city"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"mf_backup_onetime/util/custom"
	"net/http"
	"strconv"
)

type MsCityHandler struct {
	MsCityService ms_city.Service
}

func InitMsCityHandler(g *gin.Engine, ms_cityService ms_city.Service) {
	handler := &MsCityHandler{
		MsCityService: ms_cityService,
	}

	routeAPI := g.Group("/api/v1/ms_city")

	routeAPI.POST("/list", cHttp.Auth, handler.Get)
	//routeAPI.GET("/get/:id", cHttp.Auth, handler.GetByID)
	routeAPI.POST("/create", cHttp.Auth, handler.Insert)
	routeAPI.PUT("/:id", cHttp.Auth, handler.Update)
	routeAPI.DELETE("/:id", cHttp.Auth, handler.Delete)
}

// Get
// @Tags MsCity
// @Summary Get MsCity
// @Description Get MsCity
// @ID Get-MsCity
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body dto.MsCityRequestDto true "body data"
// @Success 200 {object} dto.MsCityResponseDto
// @Router /v1/ms_city/list [post]
func (h *MsCityHandler) Get(c *gin.Context) {
	var input dto.MsCityRequestDto
	var errRes dto.MsCityResponseDto

	if err := c.ShouldBindJSON(&input); err != nil {
		errRes.Code = http.StatusBadRequest
		errRes.Message = custom.ResponseMessageFailed("get", fmt.Errorf("request invalid"))
		c.JSON(http.StatusBadRequest, errRes)
		log.Errorln(err)
		return
	}

	res, err := h.MsCityService.GetList(input)
	if err != nil {
		errRes.Code = http.StatusInternalServerError
		errRes.Message = custom.ResponseMessageFailed("get", err)
		c.JSON(http.StatusInternalServerError, errRes)
		log.Errorln(err)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// Insert
// @Tags MsCity
// @Summary Insert MsCity
// @Description Insert MsCity
// @ID Insert-MsCity
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body dto.CreateMsCityDto true "body data"
// @Success 200
// @Router /v1/ms_city/create [post]
func (h *MsCityHandler) Insert(c *gin.Context) {
	var input dto.CreateMsCityDto
	var errRes dto.MsCityResponseDto

	if err := c.ShouldBind(&input); err != nil {
		errRes.Code = http.StatusBadRequest
		errRes.Message = custom.ResponseMessageFailed("insert", fmt.Errorf("request invalid"))
		c.JSON(http.StatusBadRequest, errRes)
		log.Errorln(err)
		return
	}

	res, err := h.MsCityService.Insert(c, input)
	if err != nil {
		errRes.Code = http.StatusInternalServerError
		errRes.Message = custom.ResponseMessageFailed("insert", err)
		c.JSON(http.StatusInternalServerError, errRes)
		log.Errorln(err)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// Update
// @Tags MsCity
// @Summary Update MsCity
// @Description Update MsCity
// @ID Update-MsCity
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body dto.UpdateMsCityDto true "body data"
// @Param id path int true "ID"
// @Success 200
// @Router /v1/ms_city/{id} [put]
func (h *MsCityHandler) Update(c *gin.Context) {
	var input dto.UpdateMsCityDto
	var errRes dto.MsCityResponseDto

	idS := c.Param("id")
	if idS == "" {
		err := fmt.Errorf("id invalid")
		errRes.Code = http.StatusBadRequest
		errRes.Message = custom.ResponseMessageFailed("update", err)
		c.JSON(http.StatusBadRequest, errRes)
		log.Errorln(err)
		return
	}

	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		err = fmt.Errorf("id invalid")
		errRes.Code = http.StatusBadRequest
		errRes.Message = custom.ResponseMessageFailed("update", err)
		c.JSON(http.StatusBadRequest, errRes)
		log.Errorln(err)
		return
	}

	if err := c.ShouldBind(&input); err != nil {
		errRes.Code = http.StatusBadRequest
		errRes.Message = custom.ResponseMessageFailed("update", fmt.Errorf("request invalid"))
		c.JSON(http.StatusBadRequest, errRes)
		log.Errorln(err)
		return
	}

	input.ID = int(id)
	res, err := h.MsCityService.Update(input)
	if err != nil {
		errRes.Code = http.StatusInternalServerError
		errRes.Message = custom.ResponseMessageFailed("update", err)
		c.JSON(http.StatusInternalServerError, errRes)
		log.Errorln(err)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// Delete
// @Tags MsCity
// @Summary Delete MsCity
// @Description Delete MsCity
// @ID Delete-MsCity
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path int true "ms city ID"
// @Success 200 {object} dto.MsCityResponseDto
// @Router /v1/ms_city/{id} [delete]
func (h *MsCityHandler) Delete(c *gin.Context) {
	var errRes dto.MsCityResponseDto

	idS := c.Param("id")
	if idS == "" {
		err := fmt.Errorf("id invalid")
		errRes.Code = http.StatusBadRequest
		errRes.Message = custom.ResponseMessageFailed("delete", err)
		c.JSON(http.StatusBadRequest, errRes)
		log.Errorln(err)
		return
	}

	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		err = fmt.Errorf("id invalid")
		errRes.Code = http.StatusBadRequest
		errRes.Message = custom.ResponseMessageFailed("delete", err)
		c.JSON(http.StatusBadRequest, errRes)
		log.Errorln(err)
		return
	}

	fmt.Println("id delete ", id)

	res, err := h.MsCityService.Delete(int(id))
	if err != nil {
		errRes.Code = http.StatusInternalServerError
		errRes.Message = custom.ResponseMessageFailed("delete", err)
		c.JSON(http.StatusInternalServerError, errRes)
		log.Errorln(err)
	} else {
		c.JSON(http.StatusOK, res)
	}
}
