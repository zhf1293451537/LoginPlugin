package server

import (
	api "WorkOrder/api"
	"WorkOrder/conf"
	"WorkOrder/middlewares"
	router "WorkOrder/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

var src = map[string]func(*gin.Engine) *gin.RouterGroup{
	"session": router.SessionRouter,
	"cookie":  router.CookieRouter,
	"token":   router.TokenRouter,
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	// r.Static("/static", "./static")
	// r.GET("/test", api.Test)
	//登陆界面
	r.GET("/login", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "login.html", "") })
	//注册界面
	r.GET("/regis", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "regis.html", "") })
	//账号注册
	r.POST("/regis", api.UserRegis)
	//验证码放在session中
	v1 := r.Group("/auth").Use(middlewares.Session("topgoer"))
	{
		v1.GET("/captcha", api.AuthPicture)
		v1.GET("/pic", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "auth.html", nil) })
		v1.GET("/captcha/verify/:value", func(ctx *gin.Context) {
			value, _ := ctx.Params.Get("value")
			if api.CaptchaVerify(ctx, value) {
				ctx.JSON(http.StatusOK, gin.H{"status": 0, "msg": "success"})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"status": 0, "msg": "failed"})
			}
		})
	}

	//获取对应login方式的RouterGroup
	_ = src[conf.LoginType](r)
	{

	}
	return r
}
