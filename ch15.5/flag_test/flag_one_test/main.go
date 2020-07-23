package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome Show Args")
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("arg[%d] = %v\n", index, arg)
		}
	}
}
