package service

import (
	"mf_backup_onetime/dto"
	"mf_backup_onetime/module/ms_destination"
	"mf_backup_onetime/util/custom"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type MSDestinationService struct {
	MsCityRepository ms_destination.Repository
}

func InitMSDestinationService(MsCityRepository ms_destination.Repository) ms_destination.Service {
	return &MSDestinationService{
		MsCityRepository: MsCityRepository,
	}
}

func (s *MSDestinationService) GetList(req dto.MSDestinationRequestDto) (res *dto.MSDestinationResponseDto, err error) {
	log.Info("MsCityService - GetList() - starting...")
	//if err = req.Validate(); err != nil {
	//	return res, err
	//}

	result, count, err := s.MsCityRepository.GetList(req)
	if err != nil {
		return res, err
	}

	res = new(dto.MSDestinationResponseDto)
	res.Code = http.StatusOK
	res.Message = custom.ResponseMessageSuccess("get")
	res.Items = result
	res.Count = count

	log.Info("MsCityService - GetList() - finished.")
	return res, nil
}

//func (s *MSDestinationService) Insert(ctx context.Context, req dto.CreateMsCityDto) (res *dto.MsCityResponseDto, err error) {
//	log.Info("MsCityService - Insert() - starting...")
//	if err = req.Validate(); err != nil {
//		return res, err
//	}
//
//	dataInsert := model.MsCity{
//		Name:      req.Name,
//		Longitude: req.Longitude,
//		Latitude:  req.Latitude,
//		MSAreaId:  req.MSAreaId,
//	}
//	dataInsert.InitAudit(constant.OPERATION_SQL_INSERT, "system")
//
//	err = s.MsCityRepository.Insert(&dataInsert)
//	if err != nil {
//		return res, err
//	}
//
//	res = new(dto.MsCityResponseDto)
//	res.Code = http.StatusOK
//	res.Message = custom.ResponseMessageSuccess("insert")
//	res.Item = &dataInsert
//
//	return res, nil
//}
//
//func (s *MSDestinationService) Update(req dto.UpdateMsCityDto) (res *dto.MsCityResponseDto, err error) {
//	log.Info("VersionService - Update() - starting...")
//	if err = req.Validate(); err != nil {
//		return res, err
//	}
//
//	dataUpdate := model.MsCity{
//		Name:      req.Name,
//		Longitude: req.Longitude,
//		Latitude:  req.Latitude,
//		MSAreaId:  req.MSAreaId,
//	}
//	dataUpdate.ID = req.ID
//	dataUpdate.InitAudit(constant.OPERATION_SQL_UPDATE, "system")
//
//	err = s.MsCityRepository.Update(&dataUpdate)
//	if err != nil {
//		return res, err
//	}
//
//	res = new(dto.MsCityResponseDto)
//	res.Code = http.StatusOK
//	res.Message = custom.ResponseMessageSuccess("update")
//	res.Item = &dataUpdate
//
//	return res, nil
//}
//
//func (s *MSDestinationService) Delete(id int) (res *dto.MsCityResponseDto, err error) {
//	log.Info("MsCityService - delete() - starting...")
//	if id < 1 {
//		return res, fmt.Errorf("id required")
//	}
//
//	err = s.MsCityRepository.Delete(id)
//	if err != nil {
//		return res, err
//	}
//
//	res = new(dto.MsCityResponseDto)
//	res.Code = http.StatusOK
//	res.Message = custom.ResponseMessageSuccess("delete")
//
//	return res, nil
//}
