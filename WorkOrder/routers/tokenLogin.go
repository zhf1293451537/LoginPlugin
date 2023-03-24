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
		u.GET("/articles/:id/edit", func(c *gin.Context) {
			id := c.Param("id")
			article := "nil" + id //models.GetArticleByID(id)
			// if err != nil {
			// 	c.AbortWithError(http.StatusInternalServerError, err)
			// 	return
			// }
			c.HTML(http.StatusOK, "edit_article.html", gin.H{
				"article": article,
			})
		})
		u.POST("/articles/:id/edit", func(c *gin.Context) {
			// id := c.Param("id")
			// title := c.PostForm("title")
			// content := c.PostForm("content")
			// author := c.PostForm("author")
			// article := models.Article{
			// 	ID: id,
			// 	Title: title,
			// 	Content: content,
			// 	Author: author,
			// }
			// err := article.Update()
			// if err != nil {
			// 	c.AbortWithError(http.StatusInternalServerError, err)
			// 	return
			// }
			c.Redirect(http.StatusFound, "/articles")
		})
		u.POST("/articles/:id/delete", func(c *gin.Context) {
			// id := c.Param("id")
			// // err := models.DeleteArticle(id)
			// if err != nil {
			// 	c.AbortWithError(http.StatusInternalServerError, err)
			// 	return
			// }
			c.Redirect(http.StatusFound, "/articles")
		})
		/*
		* 博客文章的发布、编辑、删除、查看等功能

		* 博客文章的分类和标签管理
		* 博客文章的搜索功能
		* 博客文章的评论功能
		* 博客文章的点赞功能
		* 管理员后台管理功能
		 */

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
