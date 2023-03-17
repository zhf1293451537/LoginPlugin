package server

import (
	api "WorkOrder/api"
	"WorkOrder/middlewares"

	"github.com/gin-gonic/gin"
)

func CookieRouter(r *gin.Engine) *gin.RouterGroup {
	//使用cookie登陆
	r.POST("/user/login", api.Cookielogin)
	u := r.Group("/v1", middlewares.AuthCookie())
	{
		u.GET("home", func(ctx *gin.Context) { ctx.String(200, "cookie auth success") })
	}
	return u
}
