package api

import (
	"WorkOrder/middlewares"
	"WorkOrder/models"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
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
		Likes:       0,
		Views:       0,
		PublishDate: time.Date(2023, 2, 4, 0, 0, 0, 0, time.UTC),
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

// 博客查看
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
	c.HTML(http.StatusOK, "show_article.html", gin.H{
		"article": article,
		"name":    cataname,
	})
}

// 博客编辑页面
func ArtEdit(c *gin.Context) {
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

// 点赞
func ArtLike(c *gin.Context) {
	articleID := c.Param("id")
	//获取username
	token := c.Request.Header.Get("jwt")
	myinfo, err := middlewares.ParseToken(token)
	if err != nil {
		c.JSON(500, gin.H{
			"err": err,
			"msg": "token parse error",
		})
		return
	}
	username := myinfo.UserName
	//文章不存在时是否需要返回id不存在错误？ 文章不存在，列表中也不会显示，也不会用到点赞等功能 所以暂时不需要

	//只有没点过赞的情况下才可以点赞
	userlike := &models.UserLike{}
	err = models.DB.Table("user_likes").Where("user_name = ?", username).Where("article_id = ?", articleID).First(userlike).Error
	if err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "user has likes"})
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "database error"})
		return
	}
	//没点过赞才可以执行以下
	err = models.DB.Table("articles").Where("id = ?", articleID).UpdateColumn("likes", gorm.Expr("likes + ?", 1)).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "database error"})
		return
	}
	uintVal, err := strconv.ParseUint(articleID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "uintVal error"})
		return
	}
	// 在user_likes中创建记录
	ul := &models.UserLike{
		UserName:  username,
		ArticleID: uint(uintVal),
	}
	models.DB.Create(ul)
	c.JSON(http.StatusOK, gin.H{"msg": "likes + 1 success"})
}

// 取消点赞
func DisArtLike(c *gin.Context) {
	articleID := c.Param("id")
	//获取username
	token := c.Request.Header.Get("jwt")
	myinfo, err := middlewares.ParseToken(token)
	if err != nil {
		c.JSON(500, gin.H{
			"err": err,
			"msg": "token parse error",
		})
		return
	}
	username := myinfo.UserName
	//只有点过赞才能取消点赞
	userlike := &models.UserLike{}
	err = models.DB.Table("user_likes").Where("user_name = ?", username).Where("article_id = ?", articleID).First(userlike).Error
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"err": "user not likes"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "database error"})
		return
	}
	//文章点赞量是否可以小于0？只有点赞后才可以取消点赞
	err = models.DB.Table("articles").Where("id = ?", articleID).UpdateColumn("likes", gorm.Expr("likes - ?", 1)).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "database error"})
		return
	}
	//删除user_likes 表中的记录
	err = models.DB.Table("user_likes").Where("user_name = ?", username).Where("article_id = ?", articleID).Delete(&models.UserLike{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "likes - 1 success"})
}

// 获取文章归档
func GetArchive(c *gin.Context) {
	articles, err := models.GetArticleByArchive()
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "list get error",
			"err": err,
		})
		// c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	//对日期进行排序
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].PublishDate.After(articles[j].PublishDate)
	})
	archiveMap := make(map[string][]*models.Article)
	for _, article := range articles {
		key := article.PublishDate.Format("2006-01")
		archiveMap[key] = append(archiveMap[key], article)
	}
	c.HTML(http.StatusOK, "archives.html", gin.H{"ArchiveMap": archiveMap})
}

func ArtPostByTime(c *gin.Context) {
	// 从表单中获取文章数据
	title := c.PostForm("title")
	content := c.PostForm("content")
	author := c.PostForm("author")
	timestampStr := c.PostForm("publishtime")
	log.Println(timestampStr)
	if timestampStr == "" {
		//没有设定时间时
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Not Set Publish Time",
		})
		return
	}
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		// 处理错误
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Publish Time convert int64 error",
		})
		return
	}

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
		Likes:       0,
		PublishDate: time.Unix(0, timestamp*int64(time.Millisecond)),
	}

	//将article加入队列
	//可以使用time.AfterFunc 计时器但是高并发时消耗资源过多
	/*
		或者
		可以使用一个定时器来每秒钟触发一次，然后在定时器触发时检查需要执行的任务并执行它们。
		这样可以减少计时器堆的大小，并且可以更好地控制任务的执行时间。
	*/
	log.Println(article)
	time.AfterFunc(time.Until(article.PublishDate), func() {
		err := article.Create()
		if err != nil {
			c.JSON(500, gin.H{
				"err": err,
				"msg": "article create error",
			})
			return
		}
	})
	// log.Println(article)
	// 重定向到文章列表页面
	c.JSON(200, gin.H{
		"msg": "article by time create success",
	})
	log.Println("重定向到文章列表页面")
	// c.Redirect(http.StatusFound, "/articles")
}

// articles rank by star
func ArtRank(c *gin.Context) {
	articles, err := models.GetArticleListByLike()
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "list get error",
			"err": err,
		})
		// c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	sort.Slice(articles, func(i, j int) bool { return articles[i].Likes > articles[j].Likes })
	c.HTML(http.StatusOK, "article_rank_list.html", gin.H{
		"articles": articles,
	})
}

//阅读统计功能
func GetReadTime(c *gin.Context) {
	articles, err := models.GetArticleByViews()
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "list get error",
			"err": err,
		})
		// c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	sort.Slice(articles, func(i, j int) bool { return articles[i].Views > articles[j].Views })
	c.HTML(http.StatusOK, "article_read.html", gin.H{
		"articles": articles,
	})
}

//相似文章推荐
func SimilarArticlesHandler(c *gin.Context) {
	articleID := c.Param("id")
	_, cataid, err := models.GetTitleCataidById(articleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	articles := &[]models.Article{}
	if err := models.DB.Where("cataid = ?", cataid).Find(articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "show_similar_article.html", gin.H{"art": articles})
}

//获取用户浏览记录
func GetRecord(c *gin.Context) {
	token := c.Request.Header.Get("jwt")
	MyClaim, err := middlewares.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Token is fail"})
		return
	}
	userID := MyClaim.UserName
	recordList, err := models.GetRecordList(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "history database is fail"})
		return
	}
	sort.Slice(recordList, func(i, j int) bool { return recordList[i].RecordTime.After(recordList[j].RecordTime) })
	c.HTML(http.StatusOK, "show_record.html", gin.H{"history": recordList})
}
