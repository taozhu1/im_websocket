package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name    string // 群名
	OwnerId uint   // 群主
	Icon    string // 群头像
	Type    int    // 类型：普通群 vip群
	Desc    string
}

func (table *Group) TableName() string {
	return "group"
}
