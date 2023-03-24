package api

import (
	"WorkOrder/middlewares"
	"WorkOrder/models"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAvatar(c *gin.Context) {
	token := c.Request.Header.Get("jwt")
	myclaim, err := middlewares.ParseToken(token)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{"err": err.Error()})
		c.Abort()
		return
	} else {
		username := myclaim.UserName
		result, err := models.RedisClient.HMGet(username, "avatar").Result()
		if err != nil {
			c.JSON(400, gin.H{"msg": "database error" + err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"msg":    "查询成功",
			"avatar": "/Users/zhf/Desktop/images" + result[0].(string),
		})
	}
}
func SetAvatar(c *gin.Context) {
	token := c.Request.Header.Get("jwt")
	myclaim, err := middlewares.ParseToken(token)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{"err": err.Error()})
		c.Abort()
		return
	} else {
		//获取头像文件
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"msg": "file get error" + err.Error()})
			return
		}
		defer file.Close()
		fp, err := avatarsave(file, header.Filename)
		if err != nil {
			c.JSON(400, gin.H{"msg": "avatar save error" + err.Error()})
			return
		}
		username := myclaim.UserName
		err = models.RedisClient.HSet(username, "avatar", fp).Err()
		if err != nil {
			c.JSON(400, gin.H{"msg": "database error" + err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"msg": "头像保存成功",
		})
	}
}
func GetUserInfo(c *gin.Context) {
	token := c.Request.Header.Get("jwt")
	myclaim, err := middlewares.ParseToken(token)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{"err": err.Error()})
		c.Abort()
		return
	} else {
		username := myclaim.UserName
		result, err := models.RedisClient.HMGet(username, "nickname", "avatar", "profile").Result()
		if err != nil {
			c.JSON(400, gin.H{"msg": "database error" + err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"msg":      "查询成功",
			"nickname": result[0],
			"avatar":   result[1],
			"profile":  result[2],
		})
	}
}
func avatarsave(f multipart.File, name string) (string, error) {
	filename := time.Now().String() + name
	log.Println(filename)
	out, err := os.Create(filepath.Join("/Users/zhf/Desktop/images", filename))
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, f)
	if err != nil {
		return "", err
	}
	return filename, nil
}
