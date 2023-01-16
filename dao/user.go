package dao

import (
	"fmt"
	"gorm.io/gorm"
	"im-websocket/model"
	"im-websocket/pkg/utils"
	"time"
)

func GetUserList() []*model.User {
	data := make([]*model.User, 10)
	DB.Find(&data)
	return data
}

func CreateUser(user model.User) *gorm.DB {
	return DB.Create(&user)
}

func DeleteUser(user model.User) *gorm.DB {
	return DB.Delete(&user)
}

func UpdateUser(user model.User) *gorm.DB {
	return DB.Model(&user).Updates(model.User{
		Name:     user.Name,
		Password: user.Password,
	})
}

func FindUserByNameAndPwd(username, password string) model.User {
	user := model.User{}
	DB.Where("name = ? and password = ?", username, password).First(&user)
	token := fmt.Sprintf("%d", time.Now().Unix())
	token = utils.MD5Encode(token)
	DB.Model(&user).Where("id = ?", user.ID).Update("identity", token)
	DB.Where("name = ? and password = ?", username, password).First(&user)
	return user
}

func FindUserByName(username string) model.User {
	user := model.User{}
	DB.Where("name = ?", username).First(&user)
	return user
}
