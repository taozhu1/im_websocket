package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"im-websocket/model"
	"im-websocket/service"
	"net/http"
)

var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GetWs(c *gin.Context) *websocket.Conn {
	// 跨域请求
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return ws
}

// WsHandler
// @Summary websocket消息
// @Schemes
// @Description websocket消息
// @Tags user
// @Accept json
// @Produce json
// @Param request body vo.AddUser true "用户信息"
// @Success 200 {string} json{"code", "msg"}
// @Router /user/login [post]
func WsHandler(c *gin.Context) {
	// 跨域请求
	conn := GetWs(c)
	// ws资源关闭
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(conn)

	uid := c.Query("uid") // 自己的id
	toUid := c.Query("toUid")

	//// 创建一个用户实例
	client := &model.Client{
		FromId:   uid,
		TargetId: toUid,
		Socket:   conn,
		Send:     make(chan []byte),
	}
	// 用户注册到用户管理上
	model.Manager.Register <- client

	go service.WSWrite(conn, "test")
	go service.WSRead(conn)
}
