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
