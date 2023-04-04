package middlewares

import (
	"WorkOrder/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RecordHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("jwt")
		MyClaim, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": "Token is fail"})
			return
		}
		articleID := c.Param("id")
		createTime := time.Now()
		title, err := models.GetArtTitleById(articleID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": "articles database is fail"})
			return
		}
		record := models.History{
			UserID:     MyClaim.UserName,
			ArticleID:  articleID,
			Title:      title,
			RecordTime: createTime,
		}
		err = record.Create()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": "record history database error"})
			return
		}
		c.Next()
	}
}
