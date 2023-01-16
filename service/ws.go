package service

import (
	"fmt"
	"github.com/gorilla/websocket"
	"im-websocket/cache"
	"time"
)

func WSWrite(ws *websocket.Conn, msg string) {
	err := cache.Publish(cache.PublishKey, msg)
	if err != nil {
		fmt.Println(err)
	}
	tm := time.Now().Format("2006-01-02 15:04:05")
	m := fmt.Sprintf("[ws][%s][publish]:%s", tm, msg)
	err = ws.WriteMessage(1, []byte(m))
	if err != nil {
		fmt.Println(err)
	}
}

func WSRead(ws *websocket.Conn) {
	msg, err := cache.Subscribe(cache.PublishKey)
	if err != nil {
		fmt.Println(err)
	}
	tm := time.Now().Format("2006-01-02 15:04:05")
	m := fmt.Sprintf("[ws][%s][subscribe]:%s", tm, msg)
	err = ws.WriteMessage(1, []byte(m))
	if err != nil {
		fmt.Println(err)
	}
}
