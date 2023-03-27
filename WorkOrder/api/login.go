package api

import (
	"WorkOrder/middlewares"
	"WorkOrder/models"
	"WorkOrder/services"
	"WorkOrder/utils"
	"log"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserRegis(c *gin.Context) {
	info := services.RegisInfo{}
	err := c.ShouldBind(&info)
	if err != nil {
		log.Println("data bind error", err)
		c.JSON(400, err)
		return
	}
	checkexists := info.CheckUsernameExists()
	if checkexists {
		log.Println("账户名已存在")
		c.JSON(400, gin.H{
			"msg": "账户名已存在",
		})
		return
	}
	pwd, err := utils.EncryptPassword(info.Password)
	if err != nil {
		log.Println("密码加密出错")
		c.JSON(400, gin.H{
			"msg": "密码加密出错",
		})
		return
	}
	err = info.RegisUser(pwd)
	if err != nil {
		log.Println("User database set error", err)
		c.JSON(400, gin.H{
			"msg": "user database set error",
		})
		return
	}
	log.Println("用户注册成功")
	c.JSON(200, gin.H{
		"msg": "user register success",
	})
}
func Cookielogin(c *gin.Context) {
	//前段md5加密密码发送后端
	//后段bcrypt加盐存储密码到数据库
	//使用结构体bindJSON数据
	info := services.UserLoginInfo{}
	err := c.ShouldBind(&info)
	if err != nil {
		log.Println(err)
		c.JSON(401, err)
		return
	}
	res, flag := info.Userverification()
	if flag {
		name, err := utils.EncryptPassword(info.Username)
		if err != nil {
			log.Println(err.Error())
		}
		err = models.RedisClient.Set(name, "YES", time.Second*120).Err()
		if err != nil {
			log.Println("redis set error", err)
			c.Abort() //redis错误终止程序 其实应该重定向
		}
		c.SetCookie("name", name, 60, "", "", false, true) //先设置cookie再设置JSON返回值否则不能发送cookie
		c.JSON(200, res)
		return
	}
	c.JSON(401, res)
}

func UserSessionlogin(c *gin.Context) {
	info := services.UserLoginInfo{}
	err := c.ShouldBind(&info)
	if err != nil {
		log.Println(err)
		c.JSON(401, err)
		return
	}
	res, flag := info.Userverification()
	if flag {
		session := sessions.Default(c)
		session.Set("sessionID", info.Username)
		session.Save()
		c.JSON(200, res)
		c.String(200, "session 设置成功")
		return
	}
	c.JSON(401, res)
}

func UserTokenlogin(c *gin.Context) {
	info := services.UserLoginInfo{}
	user := &models.UserInfo{}
	err := c.ShouldBind(&info)
	if err != nil {
		log.Println(err)
		c.JSON(401, err)
		return
	}
	log.Println(info.Password, info.Username)
	res, flag := info.Userverification()
	if flag {
		token, err := middlewares.GenerateToken(info.Username)
		if err != nil {
			c.JSON(401, gin.H{
				"msg": "token生成发生错误",
				"err": err.Error(),
			})
			return
		}
		result := models.DB.Where("username = ?", info.Username).Select("nick_name,avatar,profile").Find(&user)
		if result.Error != nil {
			c.JSON(401, gin.H{
				"msg": "database error",
				"err": err.Error(),
			})
			return
		}
		//用户信息存入缓存
		models.RedisClient.HMSet(info.Username, map[string]interface{}{
			"nickname": user.NickName,
			"avatar":   user.Avatar,
			"profile":  user.Profile,
		})
		c.Header("jwt", token)
		c.JSON(200, res)
		return
	}
	c.JSON(401, res)
}
