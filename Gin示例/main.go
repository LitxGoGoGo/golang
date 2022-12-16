package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	ginS := gin.Default()
	ginS.LoadHTMLGlob("template/*")
	ginS.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"状态": "成功",
		})
	})

	ginS.GET("/query/info", func(c *gin.Context) {
		userId := c.Query("userid")
		name := c.Query("name")

		c.JSON(200, gin.H{
			"你的名字": name,
			"你的ID": userId,
		})
	})

	ginS.GET("/params/:msg", func(c *gin.Context) {
		msg := c.Param("msg")
		fmt.Println(msg)
		c.JSON(200, gin.H{
			"msg": msg,
		})
	})

	ginS.GET("/json", func(c *gin.Context) {
		b, _ := c.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(b, &m)
		c.JSON(200, m)
	})

	ginS.POST("/form", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"fileSize": c.PostForm("fileSize"),
			"fileName": c.PostForm("fileName"),
		})
	})

	ginS.GET("/test", func(c *gin.Context) {
		c.Redirect(301, "/index")
	})

	//404
	ginS.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", nil)
	})

	//路由组
	userGroup := ginS.Group("/user")

	{
		userGroup.GET("/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"你的ID": c.Param("id"),
			})
		})
	}

	//ginS.Use(Authorize())
	ginS.GET("/service_with_auth", Authorize(), func(c *gin.Context) {
		v1 := c.MustGet("hello").(string)
		log.Println(v1)
		c.JSON(200, gin.H{
			"成功": "通过中间件",
		})
	})
	ginS.Run(":8088")
}

// Authorize 自定义中间件
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username") // 用户名
		c.Set("hello", "dididi")
		if username == "123" {
			c.Next()
		} else {
			c.Abort()
			c.JSON(501, gin.H{
				"消息": "该用户" + username + "非法",
			})
		}
	}
}
