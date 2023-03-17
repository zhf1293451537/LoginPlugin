package middlewares

import (
	"WorkOrder/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func AuthCookie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("name")
		if err == nil { // 提取cookie成功，更新cookie
			errtemp := models.RedisClient.Exists(cookie).Err()
			if errtemp == redis.Nil {
				log.Println("cookie not exist", err)
			} else if errtemp != nil {
				log.Println("redis get error", err)
			} else {
				ctx.Next() //进行后续请求
				return
			}
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "Not authenticate",
		})
		ctx.Abort() //终止程序
	}
}
