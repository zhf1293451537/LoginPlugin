package api

import (
	"WorkOrder/middlewares"
	"WorkOrder/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	articleID := c.Param("id")
	objectIDStr := articleID + strings.Repeat("0", (24-len(articleID)))
	objectID, err := primitive.ObjectIDFromHex(objectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid article ID"})
		log.Println(err.Error())
		return
	}
	userID := fmt.Sprintf("%x", myinfo.UserName)
	userIDStr := userID + strings.Repeat("0", (24-len(userID)))
	userobjectID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid user ID"})
		return
	}
	content := c.PostForm("content")

	comment := models.Comment{
		ArticleID: objectID,
		UserID:    userobjectID,
		Content:   content,
		CreatedAt: time.Now(),
	}
	parentIDStr := c.Query("parent_id")
	var parentID *primitive.ObjectID
	if parentIDStr != "" {
		parentIDValue, err := primitive.ObjectIDFromHex(parentIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid parent ID"})
			return
		}
		parentID = &parentIDValue
		comment.ParentID = parentID
	}

	result, err := models.CommentsCollection.InsertOne(context.Background(), comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	comment.ID = result.InsertedID.(primitive.ObjectID)
	log.Println(comment)
	c.JSON(http.StatusCreated, comment)
}

func GetComments(c *gin.Context) {
	articleID := c.Param("id")
	objectIDStr := articleID + strings.Repeat("0", (24-len(articleID)))
	objectID, err := primitive.ObjectIDFromHex(objectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid article ID"})
		return
	}
	//从mongodb中查询文章所有评论
	cursor, err := models.CommentsCollection.Find(context.Background(), bson.M{"article_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	defer cursor.Close(context.Background())
	var comments []models.Comment
	for cursor.Next(context.Background()) {
		var comment models.Comment
		if err := cursor.Decode(&comment); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		comments = append(comments, comment)
	}
	c.JSON(http.StatusOK, comments)
}
