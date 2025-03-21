package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"go_learning/todolist/config"
	"go_learning/todolist/internal/service"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	fileService *service.FileService
	config      *config.Config
}

func NewFileHandler(fileService *service.FileService, cfg *config.Config) *FileHandler {
	return &FileHandler{
		fileService: fileService,
		config:      cfg,
	}
}

// Upload 处理文件上传
func (h *FileHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取上传文件失败"})
		return
	}

	// 从上下文中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	fileRecord, err := h.fileService.Upload(
		file,
		userID.(uint),
		h.config.Upload.BaseDir,
		h.config.Upload.MaxSize,
		h.config.Upload.AllowedTypes,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fileRecord)
}

// Delete 处理文件删除
func (h *FileHandler) Delete(c *gin.Context) {
	fileID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文件ID"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	if err := h.fileService.Delete(uint(fileID), userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文件删除成功"})
}

// Get 获取文件信息的处理方法
func (h *FileHandler) Get(c *gin.Context) {
	fileID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文件ID"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	file, err := h.fileService.GetFileByUser(uint(fileID), userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, file)
}

// List 获取文件列表
func (h *FileHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	files, total, err := h.fileService.ListFiles(userID.(uint), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"files":     files,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// GetFileMsg 处理获取文件信息请求 (为了兼容原有代码)
func (h *FileHandler) GetFileMsg(c *gin.Context) {
	fileID, err := strconv.ParseUint(c.Param("fileID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文件ID"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		// 如果未登录，尝试获取公开文件
		file, err := h.fileService.GetFile(int(fileID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, file)
		return
	}

	// 已登录用户获取自己的文件
	file, err := h.fileService.GetFileByUser(uint(fileID), userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, file)
}

// Download 处理文件下载请求 (为了兼容原有代码)
func (h *FileHandler) Download(c *gin.Context) {
	fileID, err := strconv.ParseUint(c.Param("fileID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文件ID"})
		return
	}

	// 获取文件信息
	file, err := h.fileService.GetFile(int(fileID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文件不存在"})
		return
	}

	// 提供文件下载
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.FileName))
	c.Header("Content-Type", "application/octet-stream")
	c.File(file.FilePath)
}
