package cache

import (
	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
	"strconv"
)

var RedisClient *redis.Client

func InitRedis(RedisAddr, Password, DB, PoolSize, MinIdleConn string) {
	db, _ := strconv.ParseUint(DB, 10, 64)
	poolSize, _ := strconv.Atoi(PoolSize)
	minIdleConn, _ := strconv.Atoi(MinIdleConn)
	client := redis.NewClient(&redis.Options{
		Addr:         RedisAddr,
		Password:     Password,
		DB:           int(db),
		MinIdleConns: poolSize,
		PoolSize:     minIdleConn,
	})
	_, err := client.Ping().Result()
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	RedisClient = client
}
