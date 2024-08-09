package ms_destination

import (
	"mf_backup_onetime/dto"
	"mf_backup_onetime/module/ms_destination/model"
)

type Repository interface {
	Insert(data *model.MSDestination) (err error)
	GetList(req dto.MSDestinationRequestDto) (res []*model.MSDestination, count int64, err error)
	Delete(id int64) (err error)
	Update(data *model.MSDestination) (err error)
	GetById(id int64) (res *model.MSDestination, err error)
	GetByIdArray(IdDestination []int64) (res []*model.MSDestination, count int64, err error)
}

type Service interface {
	GetList(req dto.MSDestinationRequestDto) (res *dto.MSDestinationResponseDto, err error)

	//Insert(ctx context.Context, req dto.CreateMsCityDto) (res *dto.MsCityResponseDto, err error)
	//
	//Update(req dto.UpdateMsCityDto) (res *dto.MsCityResponseDto, err error)
	//
	//Delete(id int) (res *dto.MsCityResponseDto, err error)
}
