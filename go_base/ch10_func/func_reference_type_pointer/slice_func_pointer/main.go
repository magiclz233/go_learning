package main

import "fmt"

// 切片不需要执行指针指定, 内部已经这样做了
func main() {
	m := make([]string, 2, 25)
	fmt.Println(m, len(m), cap(m)) //[ ] 2 25
	changeMe(m)
	fmt.Println(m, len(m), cap(m)) // [magic faker] 2 25
}

func changeMe(z []string) {
	z[0] = "magic"
	z[1] = "faker"
	fmt.Println(z)
}
