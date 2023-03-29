package conf

import (
	"WorkOrder/models"
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	Host      string
	Port      string
	LoginType string
	RedisHost string
	RedisPort string
	RedisPwd  string
)

func Init() {
	// 初始化配置文件
	Config := viper.New()
	Config.AddConfigPath("./conf")
	Config.SetConfigName("app")
	Config.SetConfigType("yaml")
	if err := Config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			log.Panicln("can't find conf.ini")
		} else {
			log.Println("conf.ini read occure")
		}
	}
	Config.WatchConfig() //监控配置文件并且热加载程序，不重启就可以加载新的配置文件
	Config.OnConfigChange(func(in fsnotify.Event) { log.Println("config file changed: ", in.Name) })
	//读取配置文件
	{
		Host = Config.GetString("server.host")
		Port = Config.GetString("server.port")
		LoginType = Config.GetString("setting.loginType")
		RedisHost = Config.GetString("redis.host")
		RedisPort = Config.GetString("redis.port")
		RedisPwd = Config.GetString("redis.pwd")
	}
	//数据库连接
	fmt.Println("进行Redis数据库连接")
	models.InitRedis(RedisHost+":"+RedisPort, RedisPwd)

	fmt.Println("用户信息数据库mysql连接")
	models.InitBlogMysql("root:123456@/userinfo?charset=utf8&parseTime=True&loc=Local")

	fmt.Println("mongoDB连接")
	models.InitMongo("mongodb://127.0.0.1:27017")
	// err := models.Database("")
	// if err != nil {
	// 	fmt.Println("数据库连接失败")
	// 	panic(err)
	// }
	// fmt.Println("数据库连接成功！")

}
