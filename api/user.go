package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"im-websocket/model"
	"im-websocket/pkg/utils"
	"im-websocket/service"
	"im-websocket/vo"
	"math/rand"
	"strconv"
)

// UserRegister
// @Summary 用户注册
// @Schemes
// @Description UserRegister
// @Tags user
// @Accept json
// @Produce json
// @Param request body vo.AddUser true "用户信息"
// @Success 200 {string} json{"code", "msg"}
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
	var body = &vo.AddUser{}
	if err := c.BindJSON(&body); err != nil {
		return
	}
	salt := fmt.Sprintf("%06d", rand.Int31())
	service.CreateUser(model.User{
		Name:     body.Username,
		Password: utils.MakePassword(body.Password, salt),
		Salt:     salt,
	})
	c.JSON(200, gin.H{
		"msg":  "success",
		"data": nil,
	})
}

// UserLogin
// @Summary 用户登录
// @Schemes
// @Description UserLogin
// @Tags user
// @Accept json
// @Produce json
// @Param request body vo.AddUser true "用户信息"
// @Success 200 {string} json{"code", "msg"}
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var body = &vo.AddUser{}
	if err := c.BindJSON(&body); err != nil {
		return
	}
	user := service.FindUserByName(body.Username)
	if user.Salt == "" && utils.ValidPassword(body.Password, user.Salt, body.Password) {
		c.JSON(200, gin.H{
			"msg":  "该用户不存在",
			"data": nil,
		})
	}
	user = service.FindUserByNameAndPwd(body.Username, utils.MakePassword(body.Password, user.Salt))
	c.JSON(200, gin.H{
		"msg":  "success",
		"data": user,
	})
}

// GetUserList
// @Summary 所有用户
// @Tags user
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	res := service.GetUserList()
	c.JSON(200, gin.H{
		"msg":  "success",
		"data": res,
	})
}

// DeleteUser
// @Summary 删除用户
// @Schemes
// @Description DeleteUser
// @Tags user
// @Accept json
// @Produce json
// @Param id query string true "用户ID"
// @Success 200 {string} json{"code", "msg"}
// @Router /user/deleteUser [delete]
func DeleteUser(c *gin.Context) {
	user := model.User{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	service.DeleteUser(user)
	c.JSON(200, gin.H{
		"msg":  "success",
		"data": user,
	})
}

// UpdateUser
// @Summary 更新用户
// @Schemes
// @Description UpdateUser
// @Tags user
// @Accept json
// @Produce json
// @Param request body vo.UpdateUser true "用户信息"
// @Success 200 {string} json{"code", "msg"}
// @Router /user/updateUser [put]
func UpdateUser(c *gin.Context) {
	var body = &vo.UpdateUser{}
	if err := c.BindJSON(&body); err != nil {
		return
	}
	user := model.User{
		Name:     body.Username,
		Password: body.Password,
	}
	id, _ := strconv.Atoi(body.ID)
	user.ID = uint(id)
	service.UpdateUser(user)
	c.JSON(200, gin.H{
		"msg":  "success",
		"data": nil,
	})
}
