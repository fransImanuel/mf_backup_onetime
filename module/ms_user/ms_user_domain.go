package ms_user

import (
	"mf_backup_onetime/dto"
	"mf_backup_onetime/module/ms_user/model"
	"mf_backup_onetime/schemas"
)

type Repository interface {
	GetList(req dto.MSUsersRequestDto) (res []*model.MSUser, count int64, err error)
	UserById(user_id int64) (*model.MSUser, schemas.SchemaDatabaseError)
}

type Service interface {
}
