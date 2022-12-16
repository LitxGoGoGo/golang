package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type BackMyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var salt string = "hallo123"

// GetToken Get /token?username=""
//get token
func GetToken(c *gin.Context) {
	username := c.Query("username")
	token := GenerateToken(username)

	//处理请求头
	//for k, v := range c.Request.Header {
	//}
	//fmt.Println("---------------------")
	//fmt.Println(c.Request.Header["User-Agent"][0])

	//返回响应体
	c.JSON(200, gin.H{
		"user-token": token,
	})
}

func GenerateToken(username string) string {
	c := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			Issuer:    "Laa",
			ExpiresAt: time.Now().Unix() + 60,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	//加密盐
	mySigningKey := []byte(salt)

	s, e := token.SignedString(mySigningKey)
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(token)
	fmt.Println(s)
	return s //s可以给前端使用了
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取请求头中的token
		userBackToken := c.Request.Header["X-Token"][0] // 用户名 注意在解析请求头的时候必须为驼峰大写
		log.Print("----------")
		log.Print(userBackToken)
		var myBackClaims BackMyClaims

		token, err := jwt.ParseWithClaims(userBackToken, &myBackClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(salt), nil
		}) //返回的这个盐要一致才能解析

		if claims, ok := token.Claims.(*BackMyClaims); ok && token.Valid {
			c.JSON(200, gin.H{
				"你的名字":  claims.Username,
				"你的授权人": claims.Issuer,
			})
			c.Next()
		} else {
			c.Abort()
			c.JSON(200, gin.H{
				"message": "token已失效请尝试重新登录或清空浏览器缓存",
			})
			fmt.Println(err)
		}
	}
}
