package controllers

import (
	"fmt"
	"github.com/LitxGoGoGo/ginAndGorm/models"
	"github.com/gin-gonic/gin"
	"path"
)

// HandelFileFromFrontEnd Post /fileUpload
// handle fileUpload
func HandelFileFromFrontEnd(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{
			"code":    "000",
			"message": "文件上传失败",
		})
	}
	//相对工程路径存储文件
	dst := path.Join("./upload", file.Filename)
	fileObject := models.UploadFile{
		Name: file.Filename,
		Path: dst,
	}
	fmt.Println(file.Filename)
	models.DB.Create(&fileObject)
	fmt.Println(dst)
	c.SaveUploadedFile(file, dst)
	c.JSON(200, gin.H{
		"code":    "001",
		"message": file.Filename + "上传成功",
		"data":    &fileObject,
	})
}
