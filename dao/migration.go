package dao

import (
	"fmt"
	"im-websocket/model"
	"os"
)

func Migration() {
	err := DB.
		Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
			&model.Message{},
			&model.Group{},
			&model.Contact{},
		)
	if err != nil {
		fmt.Print("init table fail")
		os.Exit(0)
	}
	return
}
