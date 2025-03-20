//go:build wireinject
// +build wireinject

package injector

import (
    "go_learning/todolist/internal/server"
    "go_learning/todolist/internal/service"
    "github.com/google/wire"
    "gorm.io/gorm"
)

// InitializeServer 初始化服务
func InitializeServer(db *gorm.DB) (*server.Server, error) {
    wire.Build(
        service.NewUserService,
        service.NewFileService,
        server.NewUserHandler,
        server.NewFileHandler,
        server.NewServer,
    )
    return &server.Server{}, nil
}