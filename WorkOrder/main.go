package main

import (
	"WorkOrder/conf"
	"WorkOrder/server"
)

func main() {
	//配置文件初始化以及数据库的连接
	conf.Init()

	//装载路由
	r := server.NewRouter()

	//运行
	r.Run(conf.Host + ":" + conf.Port)
}
