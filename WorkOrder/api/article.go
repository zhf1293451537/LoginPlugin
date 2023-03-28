package api

import (
	"WorkOrder/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
			"msg": "article database select error",
			"err": err,
		})
		// 	c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	cataname, err := models.GetCataByID(fmt.Sprint(article.Cataid))
	if err == gorm.ErrRecordNotFound {
		cataname = "分类已被删除请修改分类"
	} else if err != nil {
		c.JSON(500, gin.H{
			"msg": "cataname database select error",
			"err": err,
		})
		// 	c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.HTML(http.StatusOK, "edit_article.html", gin.H{
		"article": article,
		"name":    cataname,
	})
}
func ArtFix(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	content := c.PostForm("content")
	author := c.PostForm("author")
	cata := c.PostForm("catagory")
	cataid, err := models.GetCataId(cata)
	if err != nil {
		c.JSON(500, gin.H{
			"err": err,
			"msg": "cataid get fail",
		})
	}
	article := models.Article{
		Title:   title,
		Content: content,
		Author:  author,
		Cataid:  cataid,
	}
	err = article.UpdateById(id)
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

func Artsearch(c *gin.Context) {
	// 获取用户输入的搜索关键字
	keyword := c.Query("keyword")

	// 在文章信息表中查询所有包含搜索关键字的文章
	articles := &[]models.Article{}
	if err := models.DB.Where("title LIKE ?", "%"+keyword+"%").Find(articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 将搜索结果返回给用户
	c.HTML(200, "article_list.html", gin.H{
		"articles": articles,
	})
}

func GetListByCata(c *gin.Context) {
	cata := c.Param("catagory")
	result := &models.Catagory{}
	err := models.DB.Table("catagories").Where("name = ?", cata).Select("id").Find(&result).Error
	if err == gorm.ErrRecordNotFound {
		//分类已被删除
		c.JSON(500, "cata has deleted")
		return
	} else if err != nil {
		c.JSON(500, "cata database error")
		return
	}
	articles, err := models.GetArticleListByCata(fmt.Sprint(result.ID))
	if err == gorm.ErrRecordNotFound {
		log.Println("art is not found")
	} else if err != nil {
		c.JSON(500, gin.H{
			"msg": "cata list get error",
			"err": err,
		})
		// c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.HTML(http.StatusOK, "article_cata_list.html", gin.H{
		"articles": articles,
		"name":     cata,
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
	// 创建新分类
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
	log.Println("重定向到分类列表页面")
	// 重定向到分类列表页面
	c.JSON(200, gin.H{
		"msg": "cata create success",
	})
	// c.Redirect(http.StatusFound, "/cata/list")
}

func CataDelete(c *gin.Context) {
	id := c.Param("id")
	err := models.DB.Table("catagories").Where("id = ?", id).Delete(&models.Catagory{}).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err": err,
			"msg": "cata database error",
		})
	}
	c.JSON(200, gin.H{
		"msg": "cata delete success",
	})
}
func CataEdit(c *gin.Context) {
	id := c.Param("id")
	name := c.PostForm("name")
	cata := &models.Catagory{Name: name}
	err := models.DB.Table("catagories").Where("id = ?", id).Update(cata).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err": err,
			"msg": "cata database error",
		})
	}
	c.JSON(200, gin.H{
		"msg": "cata edit success",
	})

}
