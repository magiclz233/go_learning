package server

import (
	"fmt"
	"go_learning/todolist/internal/data"
	"go_learning/todolist/internal/service"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	fileService *service.FileService
}

func NewFileHandler(fileService *service.FileService) *FileHandler {
	return &FileHandler{fileService: fileService}
}

func (f *FileHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请上传文件",
		})
		return
	}
	// 获取文件名
	filename := filepath.Base(file.Filename)

	// 设置保存路径（这里保存在当前目录下的uploads文件夹）
	savePath := fmt.Sprintf("./uploads/%s", filename)

	// 创建uploads目录（如果不存在）
	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "服务器错误",
		})
		return
	}

	// 保存文件
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "文件上传失败",
		})
		return
	}
	// 保存文件信息到数据库
	fileInfo := &data.File{
		FileName: filename,
		FilePath: savePath,
		FileSize: file.Size,
		FileType: file.Header.Get("Content-Type"),
	}

	if err := f.SaveFileMsg(fileInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "文件信息保存失败",
		})
		return
	}
	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message":  "文件上传成功",
		"filename": filename,
		"size":     file.Size,
		"path":     savePath,
	})
}

func (f *FileHandler) Download(c *gin.Context) {
	// 从路径参数获取fileID并转换为uint
	fileID, err := strconv.Atoi(c.Param("fileID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的文件ID",
		})
		return
	}
	// 从数据库获取文件信息
	file, err := f.fileService.GetFile(fileID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "文件不存在",
		})
		return
	}

	// 检查文件是否存在
	if _, err := os.Stat(file.FilePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "文件不存在",
		})
		return
	}

	// 设置下载文件名
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.FileName))
	// 下载文件
	c.File(file.FilePath)
}

func (f *FileHandler) GetFileMsg(c *gin.Context) {
	// 从路径参数获取fileID并转换为整数
	fileID, err := strconv.Atoi(c.Param("fileID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的文件ID",
		})
		return
	}
	file, err := f.fileService.GetFile(fileID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "文件不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "文件信息获取成功",
		"filename": file.FileName,
		"size":     file.FileSize,
		"path":     file.FilePath,
	})
}

func (f *FileHandler) SaveFileMsg(file *data.File) error {
	return f.fileService.SaveFile(file)
}

// Delete 删除文件
func (f *FileHandler) Delete(c *gin.Context) {
	fileID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的文件ID",
		})
		return
	}

	// 从数据库获取文件信息
	file, err := f.fileService.GetFile(fileID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "文件不存在",
		})
		return
	}

	// 删除物理文件
	if err := os.Remove(file.FilePath); err != nil && !os.IsNotExist(err) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "删除文件失败",
		})
		return
	}

	// 从数据库中删除文件记录
	if err := f.fileService.DeleteFile(file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "删除文件记录失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "文件删除成功",
	})
}

// Get 获取文件信息
func (f *FileHandler) Get(c *gin.Context) {
	fileID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的文件ID",
		})
		return
	}

	file, err := f.fileService.GetFile(fileID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "文件不存在",
		})
		return
	}

	c.JSON(http.StatusOK, file)
}

// List 获取文件列表
func (f *FileHandler) List(c *gin.Context) {
	// 从查询参数获取分页信息
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// 从服务获取文件列表
	files, err := f.fileService.GetFileList(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取文件列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"files":     files,
		"page":      page,
		"page_size": pageSize,
	})
}
