package main

import "fmt"

func main() {
	age := 44
	fmt.Println(&age) // 0xc00000a0b8
	changeAge(&age)
	fmt.Println(&age) // 0xc00000a0b8
	fmt.Println(age)  // 24
}

func changeAge(z *int) {
	fmt.Println(z)  // 0xc00000a0b8
	fmt.Println(*z) // 44
	*z = 24
	fmt.Println(z)  // 0xc00000a0b8
	fmt.Println(*z) // 24
}
