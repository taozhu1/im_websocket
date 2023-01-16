package cache

const (
	PublishKey = "WS:CHAT"
)

// Publish 发布消息到Redis channel
func Publish(channel string, msg string) error {
	// cmd: redis publish
	err := RedisClient.Publish(channel, msg).Err()
	return err
}

// Subscribe 去Redis channel订阅消息
func Subscribe(channel string) (string, error) {
	// cmd: redis subscribe
	sub := RedisClient.Subscribe(channel)
	msg, err := sub.ReceiveMessage()
	return msg.Payload, err
}
