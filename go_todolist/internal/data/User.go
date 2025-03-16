package data

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
	Gender  string `json:"gender"`
	Password string `json:"password"`
	Email    string	`json:"email"`
	Phone    string	`json:"phone"`
	Address  string	`json:"address"`
}
