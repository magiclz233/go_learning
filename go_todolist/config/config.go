package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Redis      RedisConfig      `yaml:"redis"`
	PostgreSQL PostgreSQLConfig `yaml:"postgresql"`
	Upload     UploadConfig     `yaml:"upload"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type PostgreSQLConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type UploadConfig struct {
	BaseDir      string   `yaml:"base_dir"`      // 文件上传的基础目录
	MaxSize      int64    `yaml:"max_size"`      // 最大文件大小（字节）
	AllowedTypes []string `yaml:"allowed_types"` // 允许的文件类型
}

func LoadConfig() *Config {
	config := &Config{}

	// 获取项目根目录（从 cmd 目录向上一级）
	rootDir := filepath.Dir(filepath.Dir(os.Args[0]))

	// 构建配置文件的绝对路径
	configPath := filepath.Join(rootDir, "config", "config.yaml")
	log.Printf("配置文件路径: %s", configPath)

	data, err := os.ReadFile(configPath)
	if err != nil {
		// 如果第一次尝试失败，尝试使用相对于执行目录的路径
		execDir, _ := os.Getwd()
		configPath = filepath.Join(execDir, "..", "config", "config.yaml")
		log.Printf("尝试备用配置文件路径: %s", configPath)
		data, err = os.ReadFile(configPath)
		if err != nil {
			log.Fatalf("读取配置文件失败: %v", err)
		}
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
	return config
}
