package models

import (
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var RedisClient *redis.Client

func InitRedis(addr string, pwd string) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    pwd,
		DB:          0, // cookie和账号密码存储在0，sessionid存储在1，token存储在2
		PoolSize:    10,
		MaxRetries:  3,
		IdleTimeout: 10 * time.Second,
	})
	pong, err := RedisClient.Ping().Result()
	if err == redis.Nil {
		log.Println("Redis异常")
	} else if err != nil {
		log.Println("Redis失败:", err)
	} else {
		log.Println(pong, "redis连接成功")
	}
}

func InitMongoDB(connString string) error {
	db, err := gorm.Open("mongodb", connString)
	db.LogMode(true)
	if err != nil {
		return err
	}

	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	//数据迁移
	// DB.AutoMigrate()

	return nil
}
