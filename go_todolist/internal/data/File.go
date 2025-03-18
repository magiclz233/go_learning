package data

import "gorm.io/gorm"

type File struct {
	gorm.Model
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	UserID   uint   `json:"user_id"`
	FileSize int64  `json:"file_size"`
	FileType string `json:"file_type"`
}