package schemas

import (
	"mf_backup_onetime/dto"
)

type SchemaUnit struct {
	UnitCode string ` json:"unit_code,omitempty" binding:"required"`
	UnitNo   string ` json:"unit_no,omitempty" binding:"required"`
	Cluster  string ` json:"cluster,omitempty" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

type SchemaListUnit struct {
	dto.FilterBaseDto
	UnitCode string ` json:"unit_code,omitempty" binding:"-"`
	UnitNo   string ` json:"unit_no,omitempty" binding:"-"`
	Cluster  string ` json:"cluster,omitempty" binding:"-"`
	Address  string `json:"address" binding:"-"`
}

type SchemaImportTasklist struct {
	UnitCode string ` json:"unit_code,omitempty" binding:"required"`
	UnitNo   string ` json:"unit_no,omitempty" binding:"required"`
	//Cluster string ` json:"cluster,omitempty" binding:"required"`
	//Address string `json:"address" binding:"required"`
	PeriodNo int                 ` json:"period_no,omitempty"`
	Tasklist []SchemaUnitService `json:"tasklist"`
}

type SchemaUnitService struct {
	TypeId     int   `gorm:"column:type_id;" json:"type_id,omitempty"`
	TasklistId int64 `gorm:"column:tasklist_id;" json:"tasklist_id,omitempty"`
	//PrevRead int64 `gorm:"column:prev_read;" json:"prev_read,omitempty"`
	//PrevConsumption int64 `gorm:"column:prev_consumption;" json:"prev_consumption,omitempty"`
	CurrRead        int64 `gorm:"column:curr_read;" json:"curr_read,omitempty"`
	CurrConsumption int64 `gorm:"column:curr_consumption;" json:"curr_consumption,omitempty"`
	//Latitude string `gorm:"column:latitude;" json:"latitude,omitempty"`
	//Longitude string `gorm:"column:longitude;" json:"longitude,omitempty"`
	//StatusId int64 `gorm:"column:status_id;" json:"status_id,omitempty"`
}
