package main

import "fmt"

// 结构体需要指定是指针参数或者copy一份结构体过去
type customer struct {
	name string
	age  int
}

func main() {
	c := customer{name: "magic", age: 23}

	// name指针改变  main函数中数据未变
	changeMe(c)
	fmt.Println(c)
	fmt.Println(&c.name)

	// name 指针未变  main方法中数据改变
	changeMePointer(&c)
	fmt.Println(c)
	fmt.Println(&c.name)

}

// 值传递, z与main函数中c无关
func changeMe(z customer) {
	fmt.Println(z)
	fmt.Println(&z.name)
	z.name = "rookie"
	fmt.Println(z)
	fmt.Println(&z.name)
}

// 指针传递  z为main函数中c的指针
func changeMePointer(z *customer) {
	fmt.Println(z)
	fmt.Println(&z.name)
	z.name = "rookie"
	fmt.Println(z)
	fmt.Println(&z.name)
}
