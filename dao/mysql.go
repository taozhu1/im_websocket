package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitMySQL(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second, // 慢SQL阈值
				Colorful:                  true,        // 彩色日志
				IgnoreRecordNotFoundError: false,
				ParameterizedQueries:      false,
				LogLevel:                  logger.Info, // 日志等级
			})
	} else {
		ormLogger = logger.Default
	}
	// master db
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead, // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MYSQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采取删除并新建的方式，MYSQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 change 重命名列，MYSQL 8 之前的数据库和 MariaDb 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 命名策略：不要加s
		},
	})
	if err != nil {
		panic(err)
	}
	// db pool setting
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) // 活跃
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db

	// master slave setting
	resolverConfig := dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(connRead)},  // 读操作
		Replicas: []gorm.Dialector{mysql.Open(connWrite)}, // 写操作
		Policy:   dbresolver.RandomPolicy{},               // 负载均衡
	}
	_ = DB.Use(dbresolver.Register(resolverConfig))
	// 迁移数据库
	Migration()
}
