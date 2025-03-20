package main

import (
	"fmt"
	"go_learning/todolist/config"
	"go_learning/todolist/internal/data"
	"go_learning/todolist/wire/injector"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	cfg := config.LoadConfig()
    db := initDB(cfg)

   	// 迁移 schema
   	db.AutoMigrate(&data.User{}, &data.File{})

    // 使用 wire 生成的初始化函数
	server, err := injector.InitializeServer(db)
    if err != nil {
        log.Fatal(err)
    }

	r := server.InitRouter()
	r.Run(":8080")
}

func initDB(cfg *config.Config) *gorm.DB {
    var db *gorm.DB
    var err error

    if cfg.PostgreSQL.Host != "" {
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
        db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
        if err != nil {
            log.Fatal("failed to connect sqlite: ", err)
        }
        log.Println("使用 SQLite 数据库")
    }
    return db
}
