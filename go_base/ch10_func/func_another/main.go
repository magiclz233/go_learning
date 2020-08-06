package main

import "fmt"

func main() {
	func() {
		var a = 23
		fmt.Println(a)
		fmt.Println("main --> func(){}()")
	}()
}
