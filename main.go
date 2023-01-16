package main

import (
	"fmt"
	"im-websocket/conf"
	"im-websocket/router"
)

func main() {
	// init configuration
	conf.InitConfig()
	// init router
	r := router.NewRouter()
	err := r.Run(conf.HttpPort)
	if err != nil {
		fmt.Print("gin fail to start")
	}
}
