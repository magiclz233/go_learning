package main

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func main() {
	// 默认路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 限制表单上传大小 8MB，默认为32MB
	r.MaxMultipartMemory = 8 << 20

	r.POST("/uploadFile", uploadFile)
	r.POST("/uploadFiles", uploadFiles)

	if err := r.Run(":8080"); err != nil {
		log.Error(err.Error())
	}
}
