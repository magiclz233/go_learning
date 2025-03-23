package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

	// 添加 CORS 中间件
    // 添加 CORS 中间件
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},  // 允许的前端域名，根据实际情况修改
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

	userGroup := r.Group("/user")
	{
		userGroup.POST("/create", s.userHandler.CreateUser)
		userGroup.GET("/get/:id", s.userHandler.GetUserByID)
		userGroup.GET("/getUserList", s.userHandler.GetUserList)
		userGroup.PUT("/update/:id", s.userHandler.UpdateUser)
		userGroup.DELETE("/delete/:id", s.userHandler.DeleteUser)
	}

	fileGroup := r.Group("/file")
	{
		fileGroup.POST("/upload", s.fileHandler.Upload)
		fileGroup.GET("/download/:fileID", s.fileHandler.Download)
		fileGroup.GET("/info/:fileID", s.fileHandler.GetFileMsg)
	}

	return r
}