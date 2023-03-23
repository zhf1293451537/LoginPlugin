package router

import (
	api "WorkOrder/api"
	"WorkOrder/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenRouter(r *gin.Engine) *gin.RouterGroup {
	//使用JWT登陆
	r.POST("/v3/user/login", api.UserTokenlogin)
	u := r.Group("/v3", middlewares.AuthToken())
	{
		u.GET("/user/home", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "home.html", "") })

	}
	return u
}
