package models

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *gorm.DB
var RedisClient *redis.Client
var MongoClient *mongo.Client
var CommentsCollection *mongo.Collection

func InitMongo(url string) {
	//常见mongodb客户端
	MongoClient, err := mongo.NewClient(options.Client().ApplyURI(url).SetAuth(options.Credential{
		Username:   "admin",
		Password:   "123456",
		AuthSource: "admin",
	}))
	if err != nil {
		log.Fatal(err)
	}
	//连接mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = MongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//检查是否连接成功
	err = MongoClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	//获取comments集合
	CommentsCollection = MongoClient.Database("blogdb").Collection("comments")
}
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

func InitBlogMysql(connString string) error {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	if err != nil {
		return err
	}

	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 60)
	DB = db
	//数据迁移
	migration()

	return nil
}
func migration() {
	DB.AutoMigrate(&UserInfo{}, &Article{}, &Catagory{}, &UserLike{}, &History{})
}
