package main

import (
	"gorm.io/driver/sqlite"
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
	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "D42", Price: 100})
	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "D42") // find product with code D42
	// 更新数据 - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// 更新数据 - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// 删除数据 - delete product
	db.Delete(&product, 1)
}
