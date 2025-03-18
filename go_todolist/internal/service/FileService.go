package service

import (
	"go_learning/todolist/internal/data"

	"gorm.io/gorm"
)

type FileService struct {
	DB *gorm.DB
}

func NewFileService(db *gorm.DB) *FileService {
	return &FileService{DB: db}
}

func (s *FileService) GetFile(fileID uint) (*data.File, error) {
	var file data.File
	if err := s.DB.First(&file, fileID).Error; err != nil {
		return nil, err
	}
	return &file, nil
}
