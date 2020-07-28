package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("recover test")
	test()
	fmt.Println("Test completed")
}

func badCall() {
	panic("bad end")
}

// 从panic中恢复(Recover)
// 和java中的catch块类似
func test() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic: %s\r\n", err)
		}
	}()
	badCall()
	fmt.Println("After badCall")
}
