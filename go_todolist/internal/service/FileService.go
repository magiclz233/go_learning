package service

import (
	"fmt"
	"go_learning/todolist/internal/data"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"gorm.io/gorm"
)

type FileService struct {
	DB *gorm.DB
}

func NewFileService(db *gorm.DB) *FileService {
	return &FileService{DB: db}
}

func (f *FileService) GetFile(fileID int) (*data.File, error) {
	var file data.File
	if err := f.DB.First(&file, fileID).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func (f *FileService) SaveFile(file *data.File) error {
	return f.DB.Create(file).Error
}

// Upload 上传文件
func (f *FileService) Upload(file *multipart.FileHeader, userID uint, uploadDir string, maxSize int64, allowedTypes []string) (*data.File, error) {
	// 验证文件大小
	if file.Size > maxSize {
		return nil, fmt.Errorf("文件大小超过限制: %d > %d", file.Size, maxSize)
	}

	// 验证文件类型
	ext := strings.ToLower(filepath.Ext(file.Filename))
	ext = strings.TrimPrefix(ext, ".")
	if !f.isAllowedFileType(ext, allowedTypes) {
		return nil, fmt.Errorf("不支持的文件类型: %s", ext)
	}

	// 创建文件记录
	fileRecord := &data.File{
		UserID: userID,
	}

	// 上传文件
	if err := f.uploadFileToDisk(file, fileRecord, uploadDir); err != nil {
		return nil, err
	}

	// 保存到数据库
	if err := f.DB.Create(fileRecord).Error; err != nil {
		// 如果数据库保存失败，删除已上传的文件
		f.deleteFileFromDisk(fileRecord.FilePath)
		return nil, fmt.Errorf("保存文件记录失败: %v", err)
	}

	return fileRecord, nil
}

// Delete 删除文件
func (f *FileService) Delete(fileID uint, userID uint) error {
	var file data.File
	if err := f.DB.Where("id = ? AND user_id = ?", fileID, userID).First(&file).Error; err != nil {
		return fmt.Errorf("文件不存在或无权限删除: %v", err)
	}

	// 删除物理文件
	if err := f.deleteFileFromDisk(file.FilePath); err != nil {
		return err
	}

	// 删除数据库记录
	if err := f.DB.Delete(&file).Error; err != nil {
		return fmt.Errorf("删除文件记录失败: %v", err)
	}

	return nil
}

// GetFileByUser 获取用户的文件信息
func (f *FileService) GetFileByUser(fileID uint, userID uint) (*data.File, error) {
	var file data.File
	if err := f.DB.Where("id = ? AND user_id = ?", fileID, userID).First(&file).Error; err != nil {
		return nil, fmt.Errorf("文件不存在或无权限访问: %v", err)
	}
	return &file, nil
}

// ListFiles 获取用户的文件列表
func (f *FileService) ListFiles(userID uint, page, pageSize int) ([]data.File, int64, error) {
	var files []data.File
	var total int64

	offset := (page - 1) * pageSize
	if err := f.DB.Model(&data.File{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("获取文件总数失败: %v", err)
	}

	if err := f.DB.Where("user_id = ?", userID).Offset(offset).Limit(pageSize).Find(&files).Error; err != nil {
		return nil, 0, fmt.Errorf("获取文件列表失败: %v", err)
	}

	return files, total, nil
}

// uploadFileToDisk 将文件上传到磁盘
func (f *FileService) uploadFileToDisk(file *multipart.FileHeader, fileRecord *data.File, uploadDir string) error {
	// 确保上传目录存在
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return fmt.Errorf("创建上传目录失败: %v", err)
	}

	// 打开源文件
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("打开文件失败: %v", err)
	}
	defer src.Close()

	// 生成目标文件路径
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", fileRecord.ID, ext)
	dstPath := filepath.Join(uploadDir, filename)

	// 创建目标文件
	dst, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("创建目标文件失败: %v", err)
	}
	defer dst.Close()

	// 复制文件内容
	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("复制文件内容失败: %v", err)
	}

	// 更新文件信息
	fileRecord.FileName = file.Filename
	fileRecord.FilePath = dstPath
	fileRecord.FileSize = file.Size
	fileRecord.FileType = strings.TrimPrefix(ext, ".")

	return nil
}

// deleteFileFromDisk 从磁盘删除文件
func (f *FileService) deleteFileFromDisk(filePath string) error {
	if filePath != "" {
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("删除文件失败: %v", err)
		}
	}
	return nil
}

// isAllowedFileType 检查文件类型是否允许
func (f *FileService) isAllowedFileType(fileType string, allowedTypes []string) bool {
	for _, allowedType := range allowedTypes {
		if strings.ToLower(allowedType) == strings.ToLower(fileType) {
			return true
		}
	}
	return false
}

// DeleteFile 从数据库中删除文件记录
func (f *FileService) DeleteFile(file *data.File) error {
	return f.DB.Delete(file).Error
}

// GetFileList 获取文件列表，支持分页
func (f *FileService) GetFileList(page, pageSize int) ([]data.File, error) {
	var files []data.File
	offset := (page - 1) * pageSize

	if err := f.DB.Limit(pageSize).Offset(offset).Find(&files).Error; err != nil {
		return nil, err
	}

	return files, nil
}
