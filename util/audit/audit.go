package audit

import (
	"gorm.io/gorm"
	"time"
)

type FullAudit struct {
	ID          int            `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;" json:"created_at,omitempty"`
	CreatedUser string         `gorm:"column:created_user;not null;" json:"created_user,omitempty"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;" json:"updated_at,omitempty"`
	UpdatedUser string         `gorm:"column:updated_user;" json:"updated_user,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at,omitempty"`
	DeletedUser string         `gorm:"column:deleted_user;" json:"deleted_user,omitempty"`
}

type CreateAudit struct {
	ID          int       `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;" json:"created_at,omitempty"`
	CreatedUser string    `gorm:"column:created_user;not null;" json:"created_user,omitempty"`
}

type CreateUpdateAudit struct {
	ID          int       `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;" json:"created_at,omitempty"`
	CreatedUser string    `gorm:"column:created_user;not null;" json:"created_user,omitempty"`
	UpdatedAt   time.Time `gorm:"column:updated_at;" json:"updated_at,omitempty"`
	UpdatedUser string    `gorm:"column:updated_user;" json:"updated_user,omitempty"`
}

type CreateDeleteAudit struct {
	ID          int            `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;" json:"created_at,omitempty"`
	CreatedUser string         `gorm:"column:created_user;not null;" json:"created_user,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at,omitempty"`
	DeletedUser string         `gorm:"column:deleted_user;" json:"deleted_user,omitempty"`
}

type LookupAudit struct {
	ID int `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
}

type FullAuditMf struct {
	ID                 int64          `gorm:"column:Id;primaryKey;AUTO_INCREMENT" json:"id"`
	DeletedByUserId    *int64         `json:"deleted_by_user_id,omitempty" gorm:"column:DeletedByUserId"`
	DeletedTime        gorm.DeletedAt `json:"deleted_time,omitempty" gorm:"column:DeletedTime"`
	CreatedByUserName  string         `json:"created_by_user_name,omitempty" gorm:"column:CreatedByUserName"`
	CreatedTime        time.Time      `json:"created_time,omitempty" gorm:"column:CreatedTime"`
	ModifiedByUserName string         `json:"modified_by_user_name,omitempty" gorm:"column:ModifiedByUserName"`
	ModifiedTime       time.Time      `json:"modified_time,omitempty" gorm:"column:ModifiedTime"`
}

type CreateUpdateAuditMf struct {
	ID                 int64     `gorm:"column:Id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedByUserName  string    `json:"created_by_user_name,omitempty" gorm:"column:CreatedByUserName"`
	CreatedTime        time.Time `json:"created_time,omitempty" gorm:"column:CreatedTime"`
	ModifiedByUserName string    `json:"modified_by_user_name,omitempty" gorm:"column:ModifiedByUserName"`
	ModifiedTime       time.Time `json:"modified_time,omitempty" gorm:"column:ModifiedTime"`
}
