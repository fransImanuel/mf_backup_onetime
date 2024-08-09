package repository

import (
	"errors"
	"fmt"
	"mf_backup_onetime/constant"
	"mf_backup_onetime/dto"
	"mf_backup_onetime/module/ms_destination"
	"mf_backup_onetime/module/ms_destination/model"
	"mf_backup_onetime/util"
	"mf_backup_onetime/util/custom"
	"strings"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MsDestinationRepository struct {
	DBPostgres *gorm.DB
}

func InitMsDestinationRepository(DBPostgres *gorm.DB) ms_destination.Repository {
	return &MsDestinationRepository{
		DBPostgres: DBPostgres,
	}
}

func (r *MsDestinationRepository) GetById(id int64) (res *model.MSDestination, err error) {

	db := r.DBPostgres.Unscoped()

	if id > 0 {
		db = db.Where(`"Id" = ?`, id)
	}

	result := db.Preload(clause.Associations).Debug().Find(&res)
	if result.Error != nil {
		return res, result.Error
	}

	return res, nil
}

func (r *MsDestinationRepository) GetList(req dto.MSDestinationRequestDto) (res []*model.MSDestination, count int64, err error) {
	log.Debug("MsDestinationRepository - GetList() - starting...")

	db := r.DBPostgres.Where(`"Id" IS NOT NULL`).Debug()

	if req.SearchText != "" {
		req.SearchText = fmt.Sprintf("%%%v%%", strings.ToLower(req.SearchText))
		db = db.Where("(LOWER(name) LIKE ?)", req.SearchText)
	}
	//if req.MSAreaId > 0 {
	//	db = db.Where(`"MSAreaId" = ?`, req.MSAreaId)
	//}
	if len(req.DestinationId) > 0 {
		db = db.Where(` "Id" IN ?`, req.DestinationId)
	}

	dbCount := db.Table(constant.TABLE_MS_DESTINATION_NAME).Where(constant.DEFAULT_QUERY_SOFT_DELETE_MF).Count(&count)
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
			orderByQuery += fmt.Sprintf(`"%v"."Id" %v`, constant.TABLE_MS_DESTINATION_NAME, orderType)

		default:
			orderByQuery += fmt.Sprintf(`"%v"."CreatedTime" %v`, constant.TABLE_MS_DESTINATION_NAME, orderType)
		}
	} else {
		orderByQuery += fmt.Sprintf(`"%v"."Id" DESC`, constant.TABLE_MS_DESTINATION_NAME)
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

	log.Debug("MsDestinationRepository - GetList() - finished.")
	return res, count, nil
}

func (r *MsDestinationRepository) Insert(data *model.MSDestination) (err error) {
	log.Info("MsCityRepository - Insert() - starting...")
	if err = data.Validate(); err != nil {
		return err
	}

	result := r.DBPostgres.Create(&data)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return custom.ErrorOperationDB(data.TableName(), "insert")
	}

	log.Info("MsCityRepository - Insert() - finished...")
	return nil
}

func (r *MsDestinationRepository) Update(data *model.MSDestination) (err error) {
	log.Info("MsCityRepository - Update() - starting...")
	if err = data.Validate(); err != nil {
		return err
	}

	dataUpdate, err := r.GetById(data.ID)
	if err != nil {
		return err
	}

	if dataUpdate.ID < 1 {
		return custom.ErrorNotFoundDB(data.TableName(), "id", data.ID)
	}

	dataUpdate.Name = data.Name
	dataUpdate.Latitude = data.Latitude
	dataUpdate.Longitude = data.Longitude
	//dataUpdate.MSAreaId = data.MSAreaId
	//dataUpdate.MSArea= nil

	//dataUpdate.InitAudit(constant.OPERATION_SQL_UPDATE, data.UpdatedUser)

	result := r.DBPostgres.Updates(&dataUpdate)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return custom.ErrorOperationDB(data.TableName(), "insert")
	}

	data = dataUpdate

	log.Info("MsCityRepository - Update() - finished.")
	return nil
}

func (r *MsDestinationRepository) Delete(id int64) (err error) {
	log.Info("MsCityRepository - Delete() - starting...")

	dataDelete, err := r.GetById(id)
	if err != nil {
		return err
	}

	if dataDelete.ID < 1 {
		return custom.ErrorNotFoundDB(constant.TABLE_MS_CITY_NAME, "id", id)
	}

	dataDelete.ID = id
	//dataDelete.DeletedUser = "system"
	//
	////dataDelete.DeletedUser = username
	//dataDelete.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	result := r.DBPostgres.Save(&dataDelete)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return custom.ErrorOperationDB(constant.TABLE_MS_CITY_NAME, "delete")
	}

	log.Info("MsCityRepository - Delete() - finished.")
	return nil
}

func (r *MsDestinationRepository) GetByIdArray(IdDestination []int64) (res []*model.MSDestination, count int64, err error) {
	log.Debug("MsDestinationRepository - GetByIdArray() - starting...")

	db := r.DBPostgres

	if len(IdDestination) > 0 {
		db = db.Where(`"Id" IN ?`, IdDestination)
	}

	dbCount := db.Table(constant.TABLE_MS_DESTINATION_NAME).Where(constant.DEFAULT_QUERY_SOFT_DELETE_MF).Count(&count)
	if dbCount.Error != nil && !errors.Is(dbCount.Error, gorm.ErrRecordNotFound) {
		return res, count, err
	}

	if count < 1 {
		return res, count, nil
	}

	result := db.Preload(clause.Associations).Find(&res)
	if result.Error != nil {
		return res, count, result.Error
	}

	log.Debug("MsDestinationRepository - GetByIdArray() - finished.")
	return res, count, nil
}
