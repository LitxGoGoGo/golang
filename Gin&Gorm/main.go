package main

import (
	"github.com/LitxGoGoGo/ginAndGorm/controllers"
	"github.com/LitxGoGoGo/ginAndGorm/models"
	"github.com/gin-gonic/gin"
)

func main() {
	//gin的运行模式
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//r.Use(middleware.CORS())
	models.ConnectDatabases()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	//Token操作

	r.GET("/token", controllers.GetToken)
	r.POST("/token", controllers.Authorize())

	//文件操作
	r.POST("/fileUpload", controllers.HandelFileFromFrontEnd)
	r.Run(":8088")

}
