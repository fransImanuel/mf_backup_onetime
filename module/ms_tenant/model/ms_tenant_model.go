package model

import (
	"fmt"
	"mf_backup_onetime/constant"
	"mf_backup_onetime/util/audit"
	"time"

	"gorm.io/gorm"
)

type MSTenant struct {
	audit.FullAuditMf

	Name        string `json:"name" gorm:"column:Name"`
	Address     string `json:"address" gorm:"column:Address"`
	LogoPath    string `gorm:"Column:MSAreaId" json:"LogoPath"`
	Information string `gorm:"Column:Information" json:"information"`
	IsDeleted   *bool  `gorm:"Column:IsDeleted" json:"is_deleted"`
}

func (t *MSTenant) TableName() string {
	return constant.TABLE_MS_TENANT_NAME
}

func (t *MSTenant) Validate() error {
	if t.Name == "" {
		return fmt.Errorf("name required")
	}

	return nil
}

func (t *MSTenant) InitAudit(operation string, user string, user_id int64) {
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

func (MSTenant) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&MSTenant{})
	if err != nil {
		return err
	}

	return nil
}
