package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"im-websocket/service"
	"net/http"
)

var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WsHandler
// @Summary ws消息
// @Schemes
// @Description WsHandler
// @Tags user
// @Accept json
// @Produce json
// @Param request body vo.AddUser true "用户信息"
// @Success 200 {string} json{"code", "msg"}
// @Router /user/login [post]
func WsHandler(c *gin.Context) {
	// 跨域请求
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// ws资源关闭
	//defer func(ws *websocket.Conn) {
	//	err = ws.Close()
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//}(ws)

	go service.WSWrite(ws, "test")
	go service.WSRead(ws)
}
