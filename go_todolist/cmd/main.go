package main

import (
	"fmt"
	"go_learning/todolist/config"
	"go_learning/todolist/internal/data"
	"go_learning/todolist/internal/server"
	"go_learning/todolist/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	cfg := config.LoadConfig()

	var db *gorm.DB
	var err error

	// 判断是否有PostgreSQL配置
	if cfg.PostgreSQL.Host != "" {
		        // 使用PostgreSQL
				dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
				cfg.PostgreSQL.Host,
				cfg.PostgreSQL.Port,
				cfg.PostgreSQL.User,
				cfg.PostgreSQL.Password,
				cfg.PostgreSQL.DBName,
				cfg.PostgreSQL.SSLMode,
			)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
            log.Fatal("failed to connect postgresql: ", err)
        }
        log.Println("使用 PostgreSQL 数据库")
	} else {
		 // 使用SQLite
		 db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		 if err != nil {
			 log.Fatal("failed to connect sqlite: ", err)
		 }
		 log.Println("使用 SQLite 数据库")
	}

	// 迁移 schema
	db.AutoMigrate(&data.User{})

	userService := service.NewUserService(db)
	userHandler := server.NewUserHandler(userService)

	r := gin.Default()

	r.Group("/user")
	{
		r.POST("/create", userHandler.CreateUser)
		r.GET("/get/:id", userHandler.GetUserByID)
		r.GET("/getUserList", userHandler.GetUserList)
		r.PUT("/update/:id", userHandler.UpdateUser)
		r.DELETE("/delete/:id", userHandler.DeleteUser)
	}
	r.Run(":8080")
}