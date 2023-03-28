package api

import (
	"WorkOrder/middlewares"
	"WorkOrder/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"labix.org/v2/mgo/bson"
)

func CreateComment(c *gin.Context) {
	token := c.Request.Header.Get("jwt")
	myinfo, err := middlewares.ParseToken(token)
	if err != nil {
		c.JSON(500, gin.H{
			"err": err,
			"msg": "token parse error",
		})
		return
	}
	articleID := bson.ObjectIdHex("5f6d7f6d9c6f9a6eaf1f7b10")
	var parentID *bson.ObjectId
	parentIDStr := c.PostForm("parent_id")
	if parentIDStr != "" {
		parentIDObj := bson.ObjectIdHex(parentIDStr)
		parentID = &parentIDObj
	}
	comment := models.Comment{
		ID:        bson.NewObjectId(),
		ParentID:  parentID,
		ArticleID: articleID,
		UserID:    myinfo.UserName,
		Content:   c.PostForm("content"),
	}
	err = models.Comments.Insert(comment)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, comment)
}

func GetComments(c *gin.Context) {
	articleID := bson.ObjectIdHex(c.Param("id"))
	parentID := c.Query("parent_id")
	var query bson.M
	if parentID == "" {
		query = bson.M{"article_id": articleID, "parent_id": nil}
	} else {
		parentIDObj := bson.ObjectIdHex(parentID)
		query = bson.M{"article_id": articleID, "parent_id": &parentIDObj}
	}
	var result []models.Comment
	err := models.Comments.Find(query).All(&result)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
