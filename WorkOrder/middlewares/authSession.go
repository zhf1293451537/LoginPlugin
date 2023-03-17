package middlewares

import (
	"WorkOrder/conf"
	"log"

	"github.com/gin-contrib/sessions"
	sessionredis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func AuthSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		id := session.Get("sessionID")
		log.Println("sessionid is ", id)
		if id == nil {
			log.Println("sessionID is nil")
			ctx.String(401, "session is nil")
			ctx.Abort()
			return
		} else {
			log.Println("redis中有此sessionID")
			ctx.Next()
			return
		}
	}
}
func Session(keyPairs string) gin.HandlerFunc {
	store := SessionConfig()
	return sessions.Sessions(keyPairs, store)
}
func SessionConfig() sessions.Store {
	sessionMaxAge := 120
	sessionSecret := "mysession"
	store, err := sessionredis.NewStoreWithDB(10, "tcp", conf.RedisHost+":"+conf.RedisPort, conf.RedisPwd, "1", []byte(sessionSecret))
	if err != nil {
		log.Println("redis store error", err)
	}
	// store := cookie.NewStore([]byte(sessionSecret)) // session放在cookie中
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge,
		Path:   "/",
	})
	return store
}
