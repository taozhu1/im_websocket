package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromId   string // 发送者
	TargetId string // 接收者
	Type     int    // 消息类型：群里 私聊 广播
	Media    int    // 文本类型：文字 图片 音频
	Content  string // 消息内容
	Pic      string // 图片
	Url      string // 上传文件的访问路径
	Desc     string // 描述
	Account  int    // 其他数字统计：文件大小
}

func (table *Message) TableName() string {
	return "message"
}
