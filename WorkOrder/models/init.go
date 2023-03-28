package models

import (
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"labix.org/v2/mgo"
)

var DB *gorm.DB
var RedisClient *redis.Client
var Comments *mgo.Collection

func InitMgo() {
	session, err := mgo.Dial("mongodb://admin:123456localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	db := session.DB("blogdb")
	// users := db.C("users")
	Comments = db.C("comments")
	// // 插入新用户
	// user := User{
	// 	ID:       bson.NewObjectId(),
	// 	Name:     "Alice",
	// 	Email:    "alice@example.com",
	// 	Password: "password",
	// }
	// err = users.Insert(user)
	// if err != nil {
	// 	panic(err)
	// }

	// // 查询所有用户
	// var allUsers []User
	// err = users.Find(nil).All(&allUsers)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(allUsers)

	// // 查询指定用户
	// var oneUser User
	// err = users.Find(bson.M{"email": "alice@example.com"}).One(&oneUser)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(oneUser)

	// // 更新用户信息
	// err = users.Update(bson.M{"email": "alice@example.com"}, bson.M{"$set": bson.M{"name": "Bob"}})
	// if err != nil {
	// 	panic(err)
	// }
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
	DB.AutoMigrate(&UserInfo{}, &Article{}, &Catagory{})
}
