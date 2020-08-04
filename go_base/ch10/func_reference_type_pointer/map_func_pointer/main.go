package main

import "fmt"

// map slice chan初始化后不需要使用 pointer接收
func main() {
	m := make(map[string]int)
	changMe(m)
	fmt.Println(m["magic"])
}

func changMe(z map[string]int) {
	z["magic"] = 44
}
