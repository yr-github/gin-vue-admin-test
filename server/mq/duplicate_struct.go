package mq

import (
	"gorm.io/gorm"
	"time"
)

type Global struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type MyTask struct {
	Global
	Userid  int `json:"userid" form:"userid" gorm:"column:userid;comment:;type:int;"`
	Endid  int `json:"endid" form:"endid" gorm:"column:endid;comment:;type:int;"`
	EndDepartment  int `json:"endDepartment" form:"endDepartment" gorm:"column:end_department;comment:;type:int;"`
}
