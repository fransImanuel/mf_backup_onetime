package model

import (
	"fmt"
	"mf_backup_onetime/constant"
	"mf_backup_onetime/util/audit"
	"time"

	"gorm.io/gorm"
)

type MSDestination struct {
	audit.FullAuditMf

	Code             string `gorm:"column:Code" json:"code" `
	Name             string `gorm:"column:Name" json:"name" `
	Address          string `gorm:"column:Address" json:"address" `
	LogoPath         string `gorm:"Column:MSAreaId" json:"LogoPath"`
	Latitude         string `gorm:"Column:Latitude" json:"latitude"`
	Longitude        string `gorm:"Column:Longitude" json:"longitude"`
	Information      string `gorm:"Column:Information" json:"information"`
	TenantId         int64  `gorm:"Column:TenantId" json:"tenant_id"`
	IsDeleted        *bool  `gorm:"Column:IsDeleted" json:"is_deleted"`
	FrequencyArrival int64  `gorm:"Column:FrequencyArrival" json:"frequency_arrival"`
	MeterNumber      string `gorm:"Column:MeterNumber" json:"meter_number"`
	IsActive         *bool  `gorm:"Column:IsActive" json:"is_active"`
	MSTypeLocId      int64  `gorm:"Column:MSTypeLocId" json:"ms_type_loc_id"`
	MSCityId         int64  `gorm:"Column:MSCityId" json:"ms_city_id"`
	VendorId         int64  `gorm:"Column:vendor_id" json:"vendor_id"`
}

func (t *MSDestination) TableName() string {
	return constant.TABLE_MS_DESTINATION_NAME
}

func (t *MSDestination) Validate() error {
	if t.Name == "" {
		return fmt.Errorf("name required")
	}

	return nil
}

func (t *MSDestination) InitAudit(operation string, user string, user_id int64) {
	timeNow := time.Now()
	switch operation {
	case constant.OPERATION_SQL_INSERT:
		t.CreatedByUserName = user
		t.CreatedTime = timeNow
		t.ModifiedByUserName = user
		t.ModifiedTime = timeNow
	case constant.OPERATION_SQL_UPDATE:
		t.ModifiedByUserName = user
		t.ModifiedTime = timeNow
	case constant.OPERATION_SQL_DELETE:
		t.DeletedByUserId = &user_id
		t.DeletedTime = gorm.DeletedAt{Time: timeNow, Valid: true}
	}
}

func (MSDestination) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&MSDestination{})
	if err != nil {
		return err
	}

	return nil
}
