package middlewares

import (
	"WorkOrder/models"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const TokenExpireDuration = time.Second * 120

var Secret = []byte("Sett")

type MyClaim struct {
	UserName string
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	cla := MyClaim{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "lx-jwt",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	return token.SignedString(Secret)
}
func ParseToken(tokenString string) (*MyClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaim{}, func(t *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
func AuthToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("jwt")
		myclaim, err := ParseToken(token)
		if err != nil {
			log.Println(err)
			ctx.JSON(401, gin.H{"err": err.Error()})
			ctx.Abort()
			return
		} else {
			errtemp := models.RedisClient.Get(myclaim.UserName).Err()
			if errtemp != nil {
				log.Println(errtemp)
				ctx.JSON(401, gin.H{
					"msg": "token错误",
					"err": errtemp.Error(),
				})
				ctx.Abort()
				return
			}
			log.Println("username : ", myclaim.UserName)
			ctx.JSON(200, gin.H{"msg": "token auth success"})
			ctx.Next()
			return
		}
	}
}
