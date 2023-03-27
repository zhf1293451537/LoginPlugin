package api

import (
	"WorkOrder/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ArtPost(c *gin.Context) {
	// 从表单中获取文章数据
	title := c.PostForm("title")
	content := c.PostForm("content")
	author := c.PostForm("author")
	cata := c.PostForm("catagory")
	cataid, err := models.GetCataId(cata)
	if err != nil {
		c.JSON(500, gin.H{
			"err": err,
			"msg": "cataid get error",
		})
		// c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// 创建新文章
	article := models.Article{
		Title:       title,
		Content:     content,
		Author:      author,
		Cataid:      cataid,
		PublishDate: time.Now(),
	}
	err = article.Create()
	if err != nil {
		c.JSON(500, gin.H{
			"err": err,
			"msg": "article create error",
		})
		// c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// log.Println(article)
	log.Println("重定向到文章列表页面")
	// 重定向到文章列表页面
	c.JSON(200, gin.H{
		"msg": "article create success",
	})
	// c.Redirect(http.StatusFound, "/articles")
}
func ArtGet(c *gin.Context) {
	id := c.Param("id")
	article, err := models.GetArticleByID(id)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "database select error",
			"err": err,
		})
		// 	c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.HTML(http.StatusOK, "edit_article.html", gin.H{
		"article": article,
	})
}
func AriFix(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	content := c.PostForm("content")
	author := c.PostForm("author")
	article := models.Article{
		Title:   title,
		Content: content,
		Author:  author,
	}
	err := article.UpdateById(id)
	if err != nil {
		log.Println("articles fix success")
		// c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, gin.H{
		"msg": "article fix success",
	})
	// c.Redirect(http.StatusFound, "/articles")
}
func ArtDelete(c *gin.Context) {
	id := c.Param("id")
	err := models.DeleteArticle(id)
	if err != nil {
		log.Println("articles delete success")
		// c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, gin.H{
		"msg": "article delete success",
	})
	// c.Redirect(http.StatusFound, "/articles")
}
func ArtList(c *gin.Context) {
	articles, err := models.GetArticleList()
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "list get error",
			"err": err,
		})
		// c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.HTML(http.StatusOK, "article_list.html", gin.H{
		"articles": articles,
	})
}

func GetListByCata(c *gin.Context) {
	cata := c.Param("catagory")
	result := &models.Catagory{}
	err := models.DB.Table("catagory").Where("name = ?", cata).Select("id").Find(&result).Error
	if err != nil {
		c.JSON(500, "cata database error")
	}
	log.Println("catagory id is :", result.ID)
	articles, err := models.GetArticleListByCata(fmt.Sprint(result.ID))
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "cata list get error",
			"err": err,
		})
		// c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.HTML(http.StatusOK, "article_list.html", gin.H{
		"articles": articles,
	})
}

func GetAllCata(c *gin.Context) {
	catas, err := models.GetCataList()
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "list get error",
			"err": err,
		})
		// c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.HTML(http.StatusOK, "cata_list.html", gin.H{
		"catas": catas,
	})
}

func CataPost(c *gin.Context) {
	// 从表单中获取分类数据
	name := c.PostForm("name")
	// 创建新文章
	cata := models.Catagory{
		Name: name,
	}
	err := cata.Create() // 添加分类不能重复功能
	if err != nil {
		c.JSON(500, gin.H{
			"err": err,
			"msg": "cata create error",
		})
		// c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	log.Println("重定向到文章列表页面")
	// 重定向到文章列表页面
	c.JSON(200, gin.H{
		"msg": "cata create success",
	})
	// c.Redirect(http.StatusFound, "/articles")
}

func CataDelete(c *gin.Context) {

}
func CataEdit(c *gin.Context) {

}
