package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func uploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传文件出错")
	}
	if file != nil {
		err = c.SaveUploadedFile(file, "E:\\"+file.Filename)
		if err != nil {
			c.String(500, "文件解析错误")
		}
		c.String(http.StatusOK, file.Filename)
	}
}

func uploadFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
	}
	if form != nil {
		// 获取所有文件
		files := form.File["files"]
		// range 遍历
		for _, file := range files {
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
			}
		}
		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	}

}
