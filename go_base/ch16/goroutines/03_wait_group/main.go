package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go one()
	go two()
	wg.Wait()
}

func one() {
	for i := 0; i < 50; i++ {
		fmt.Println("One : ", i)
		time.Sleep(1e7)
	}
	wg.Done()
}

func two() {
	for i := 0; i < 50; i++ {
		fmt.Println("Two : ", i)
		time.Sleep(5e7)

	}
	wg.Done()
}
