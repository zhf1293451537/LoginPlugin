package main

import (
	"WorkOrder/conf"
	"WorkOrder/server"
)

func main() {
	//配置文件初始化以及数据库的连接
	conf.Init()
	//
	// log.Println(time.Now().Add(time.Second * 60).Unix())
	//装载路由
	r := server.NewRouter()

	//运行
	r.Run(conf.Host + ":" + conf.Port)
}

/*
压力测试
ab -H "jwt:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyTmFtZSI6IjEyOTM0NTE1MzciLCJleHAiOjE2ODA0ODMyNjUsImlzcyI6Imx4LWp3dCJ9.SyWMYB3Viy4qXDasOrTFMuIL-zthS8Y3dsJjzInkAz8" -n 1000 -c 100 http://127.0.0.1:8080/v3/articles
*/
