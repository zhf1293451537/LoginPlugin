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
		//显示主页包括导航栏
		u.GET("/user/home", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "home.html", "") })
		//显示用户个人信息页面
		u.GET("/user/infoInter", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "info.html", "") })
		u.GET("/user/avatarinter", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "avatar.html", "") })

		//获取用户个人信息
		u.GET("/user/info", api.GetUserInfo)
		//获取用户头像
		u.GET("/user/avatar", api.GetAvatar)
		//修改用户头像
		u.POST("/user/avatar", api.SetAvatar)

		//博客文章操作
		//博客发布页面
		u.GET("/articles/create", func(c *gin.Context) { c.HTML(http.StatusOK, "create_article.html", gin.H{}) })
		//博客发布
		u.POST("/articles/create", api.ArtPost)
		//博客查看
		u.GET("/articles/:id/edit", api.ArtGet)
		//博客编辑
		u.POST("/articles/:id/edit", api.ArtFix)
		//博客删除
		u.POST("/articles/:id/delete", api.ArtDelete)
		//博客列表
		u.GET("/articles", api.ArtList)

		//博客文章的分类和标签管理
		//新增分类
		u.POST("/articles/cata/create/", api.CataPost)
		//编辑分类
		u.POST("/articles/cata/:id/edit/", api.CataEdit)
		//获取所有分类
		u.GET("/articles/cata/list", api.GetAllCata)
		//删除分类
		u.POST("/articles/cata/:id/delete", api.CataDelete)

		//获取分类下的所有文章列表
		u.GET("/articles/cata/:catagory/list", api.GetListByCata)

		//博客文章的搜索功能
		u.GET("/articles/search", api.Artsearch)
		//博客文章的评论功能

		// 查询评论列表
		u.GET("/articles/:id/comments", api.GetComments)
		// 创建新评论
		u.POST("/articles/:id/comments", api.CreateComment)

		//博客文章的点赞功能
		u.POST("/articles/:id/likes", api.ArtLike)
		u.DELETE("/articles/:id/likes", api.DisArtLike)
		//管理员后台管理功能

		/*
			2. 高级功能
			* 博客文章的分享功能
			* 博客文章的推荐功能
			* 博客文章的归档功能
			* 博客文章的订阅功能
			* 博客文章的定时发布功能
			* 博客文章的热门排行功能
			* 博客文章的阅读统计功能
			* 博客文章的相关推荐功能
			3. 可选功能
			* 第三方登录功能，例如使用GitHub、Google、微信等账号进行登录
			* 多语言支持功能，例如支持中文、英文、日文等语言
			* 主题切换功能，例如支持多种主题风格
			* 邮件通知功能，例如评论回复、订阅更新等邮件通知
		*/
	}
	return u
}
