package main

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("无法链接sqlite!")
	}

	// 自动迁移
	db.AutoMigrate(&Product{})
	fmt.Println("数据库表创建成功")

	// 创建记录
	newProduct := Product{Code: "D42", Price: 100}
	db.Create(&newProduct)
	fmt.Printf("创建商品记录: Code=%s, Price=%d\n", newProduct.Code, newProduct.Price)

	// 查询记录
	var product Product
	db.First(&product, 1)
	fmt.Printf("查询ID=1的商品: Code=%s, Price=%d\n", product.Code, product.Price)

	db.First(&product, "code = ?", "D42")
	fmt.Printf("查询Code=D42的商品: ID=%d, Price=%d\n", product.ID, product.Price)

	// 更新单个字段
	db.Model(&product).Update("Price", 200)
	fmt.Printf("更新后的价格: Price=%d\n", product.Price)

	// 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	fmt.Printf("批量更新后: Code=%s, Price=%d\n", product.Code, product.Price)

	// 删除记录
	db.Delete(&product, 1)
	fmt.Println("删除ID=1的商品记录")

	// 验证删除
	var count int64
	db.Model(&Product{}).Count(&count)
	fmt.Printf("当前数据库中商品总数: %d\n", count)
}
