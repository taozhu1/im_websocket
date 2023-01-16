package service

import (
	"gorm.io/gorm"
	"im-websocket/dao"
	"im-websocket/model"
)

func GetUserList() []*model.User {
	return dao.GetUserList()
}

func CreateUser(user model.User) *gorm.DB {
	return dao.CreateUser(user)
}

func DeleteUser(user model.User) *gorm.DB {
	return dao.DeleteUser(user)
}

func UpdateUser(user model.User) *gorm.DB {
	return dao.UpdateUser(user)
}

func FindUserByNameAndPwd(username, password string) model.User {
	return dao.FindUserByNameAndPwd(username, password)
}

func FindUserByName(username string) model.User {
	return dao.FindUserByName(username)
}
