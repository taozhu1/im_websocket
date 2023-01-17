package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	FromId   int64 // 来源用户
	TargetId int64 // 目标用户
	Type     int   // 类型：0 1 3
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}
