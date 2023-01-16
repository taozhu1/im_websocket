package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name          string
	Password      string
	Salt          string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LogoutTime    time.Time `gorm:"column:logout_time" json:"login_out_time"`
	isLogout      bool
	DeviceInfo    string
}

func (table *User) TableName() string {
	return "user"
}
