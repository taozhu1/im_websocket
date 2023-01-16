package conf

import (
	"fmt"
	logging "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"im-websocket/cache"
	"im-websocket/dao"
)

var (
	AppMode     string
	HttpPort    string
	DSN         string
	RedisAddr   string
	Password    string
	DB          string
	PoolSize    string
	MinIdleConn string
)

func InitConfig() {
	//从本地读取环境变量
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	loadConfig(file)
	loadServer()
}

// loadConfig load configuration
func loadConfig(file *ini.File) {
	LoadServer(file)
	LoadMySQL(file)
	LoadRedis(file)
	logging.Info("[load config] Server AppMode:%v HttpPort:%v", AppMode, HttpPort)
	logging.Info("[load config] MySQL DSN:%v", DSN)
	logging.Info("[load config] Redis RedisAddr:%v", RedisAddr)
}

func loadServer() {
	// load mysql
	dao.InitMySQL(DSN, DSN)
	// load redis
	cache.InitRedis(RedisAddr, Password, DB, PoolSize, MinIdleConn)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMySQL(file *ini.File) {
	DSN = file.Section("mysql").Key("DSN").String()
}

func LoadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	Password = file.Section("redis").Key("Password").String()
	DB = file.Section("redis").Key("DB").String()
	PoolSize = file.Section("redis").Key("PoolSize").String()
	MinIdleConn = file.Section("redis").Key("MinIdleConn").String()
}
