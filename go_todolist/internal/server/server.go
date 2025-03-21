package server

import "github.com/gin-gonic/gin"

type Server struct {
	userHandler *UserHandler
	fileHandler *FileHandler
}

func NewServer(userHandler *UserHandler, fileHandler *FileHandler) *Server {
	return &Server{
		userHandler: userHandler,
		fileHandler: fileHandler,
	}
}

func (s *Server) InitRouter() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/user")
	{
		userGroup.POST("/create", s.userHandler.CreateUser)
		userGroup.GET("/get/:id", s.userHandler.GetUserByID)
		userGroup.GET("/getUserList", s.userHandler.GetUserList)
		userGroup.PUT("/update/:id", s.userHandler.UpdateUser)
		userGroup.DELETE("/delete/:id", s.userHandler.DeleteUser)
	}

	// 文件相关路由
	fileGroup := r.Group("/api/files")
	// 如果有认证中间件，可以在这里添加
	// fileGroup.Use(middleware.AuthMiddleware())
	{
		// 基本文件操作
		fileGroup.POST("/upload", s.fileHandler.Upload) // 上传文件
		fileGroup.DELETE("/:id", s.fileHandler.Delete)  // 删除文件
		fileGroup.GET("/:id", s.fileHandler.Get)        // 获取文件信息
		fileGroup.GET("/list", s.fileHandler.List)      // 获取文件列表

		// 保留原有的API兼容性
		fileGroup.GET("/download/:fileID", s.fileHandler.Download) // 下载文件
		fileGroup.GET("/info/:fileID", s.fileHandler.GetFileMsg)   // 获取文件信息
	}

	return r
}
