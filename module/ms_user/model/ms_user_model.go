package model

import (
	"mf_backup_onetime/constant"
	mMSTenant "mf_backup_onetime/module/ms_tenant/model"
	"mf_backup_onetime/util/audit"
	"time"

	"gorm.io/gorm"
)

type MSUser struct {
	audit.FullAuditMf
	NIK               string              `json:"nik" gorm:"column:NIK"`
	FullName          string              `json:"full_name" gorm:"column:FullName"`
	UserName          string              `json:"user_name" gorm:"column:UserName"`
	BirthDate         time.Time           `json:"birth_date" gorm:"column:BirthDate"`
	Email             string              `json:"email" gorm:"column:Email"`
	Phone             string              `json:"phone" gorm:"column:Phone"`
	Password          string              `json:"password" gorm:"column:Password"`
	IsPasswordChanged *bool               `json:"is_password_changed" gorm:"column:IsPasswordChanged"`
	PhotoPath         string              `json:"photo_path" gorm:"column:PhotoPath"`
	Infomation        string              `json:"infomation" gorm:"column:Infomation"`
	IsActive          *bool               `json:"is_active" gorm:"column:IsActive"`
	IsResign          *bool               `json:"is_resign" gorm:"column:IsResign"`
	IsEditable        *bool               `json:"is_editable" gorm:"column:IsEditable"`
	TenantId          int64               `json:"tenant_id" gorm:"column:TenantId"`
	Tenant            *mMSTenant.MSTenant `json:"tenant" gorm:"foreignKey:TenantId"`
	IsDeleted         *bool               `gorm:"Column:IsDeleted" json:"is_deleted"`
}

func (t *MSUser) TableName() string {
	return constant.TABLE_MS_USER_NAME
}

func (t *MSUser) Validate() error {
	//if t.Name == "" {
	//	return fmt.Errorf("name required")
	//}

	return nil
}

func (t *MSUser) InitAudit(operation string, user string, user_id int64) {
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

func (MSUser) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&MSUser{})
	if err != nil {
		return err
	}

	return nil
}
