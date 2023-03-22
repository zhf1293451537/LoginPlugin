package router

import (
	api "WorkOrder/api"
	"WorkOrder/middlewares"

	"github.com/gin-gonic/gin"
)

func SessionRouter(r *gin.Engine) *gin.RouterGroup {
	var u *gin.RouterGroup
	//使用session登陆
	v2 := r.Group("/v2", middlewares.Session("sessionID"))
	{
		v2.POST("/user/login", api.UserSessionlogin)

		u := v2.Group("/user", middlewares.AuthSession())
		{
			u.GET("/home", func(ctx *gin.Context) { ctx.String(200, "session auth success") })

		}
	}
	return u
}
