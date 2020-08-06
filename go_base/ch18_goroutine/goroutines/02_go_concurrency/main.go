package main

import (
	"fmt"
	"time"
)

func main() {
	go one()
	go two()
	time.Sleep(1e9)
}

func one() {
	for i := 0; i < 50; i++ {
		fmt.Println("One : ", i)
	}
}

func two() {
	for i := 0; i < 50; i++ {
		fmt.Println("Two : ", i)
	}
}
