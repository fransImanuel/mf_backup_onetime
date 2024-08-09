package repository

import (
	"errors"
	"fmt"
	"mf_backup_onetime/constant"
	"mf_backup_onetime/dto"
	"mf_backup_onetime/module/ms_user"
	"mf_backup_onetime/module/ms_user/model"
	"mf_backup_onetime/schemas"
	"mf_backup_onetime/util"
	"strings"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MSUserRepository struct {
	DBPostgres *gorm.DB
}

func InitMSUserRepository(DBPostgres *gorm.DB) ms_user.Repository {
	return &MSUserRepository{
		DBPostgres: DBPostgres,
	}
}

func (r *MSUserRepository) GetList(req dto.MSUsersRequestDto) (res []*model.MSUser, count int64, err error) {
	log.Debug("MSUserRepository - GetList() - starting...")

	db := r.DBPostgres.Where(`"Id" IS NOT NULL`).Debug()

	if req.SearchText != "" {
		req.SearchText = fmt.Sprintf("%%%v%%", strings.ToLower(req.SearchText))
		db = db.Where("(LOWER(FullName) LIKE ?)", req.SearchText)
	}
	if len(req.UserId) > 0 {
		db = db.Where(` "Id" IN ?`, req.UserId)
	}

	dbCount := db.Table(constant.TABLE_MS_USER_NAME).Where(constant.DEFAULT_QUERY_SOFT_DELETE_MF).Count(&count)
	if dbCount.Error != nil && !errors.Is(dbCount.Error, gorm.ErrRecordNotFound) {
		return res, count, err
	}

	if count < 1 {
		return res, count, nil
	}

	orderByQuery := ""
	if req.OrderField != "" {
		orderColumn, orderType := util.SplitOrderQuery(req.OrderField)
		switch orderColumn {
		case "id":
			orderByQuery += fmt.Sprintf(`"%v"."Id" %v`, constant.TABLE_MS_USER_NAME, orderType)

		default:
			orderByQuery += fmt.Sprintf(`"%v"."CreatedTime" %v`, constant.TABLE_MS_USER_NAME, orderType)
		}
	} else {
		orderByQuery += fmt.Sprintf(`"%v"."Id" DESC`, constant.TABLE_MS_USER_NAME)
	}
	db = db.Order(orderByQuery)

	if req.FilterPage > 0 && req.FilterLimit > 0 {
		offset := req.FilterPage*req.FilterLimit - req.FilterLimit
		db = db.Limit(req.FilterLimit).Offset(offset)
	}

	result := db.Preload(clause.Associations).Find(&res)
	if result.Error != nil {
		return res, count, result.Error
	}

	log.Debug("MSUserRepository - GetList() - finished.")
	return res, count, nil
}

func (r *MSUserRepository) UserById(user_id int64) (*model.MSUser, schemas.SchemaDatabaseError) {

	var user model.MSUser
	db := r.DBPostgres.Debug()

	Find := db.Preload(clause.Associations).First(&user, user_id)

	if Find.Error != nil {
		log.Errorln("❌ Error Find User: ", Find.Error.Error())
		return nil, schemas.SchemaDatabaseError{Error: Find.Error, Code: 500, Message: "Error Find User"}
	}
	if Find.RowsAffected < 1 {
		return nil, schemas.SchemaDatabaseError{Error: fmt.Errorf("Error Find User"), Code: 500, Message: "Error Find User"}
	}

	return &user, schemas.SchemaDatabaseError{}
}
