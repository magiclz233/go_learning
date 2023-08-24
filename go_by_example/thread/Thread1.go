package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	odd := make(chan int)
	even := make(chan int)

	// 奇数线程
	go func() {
		defer wg.Done()

		for i := 1; i <= 10; i += 2 {
			fmt.Println("odd：", i)
			odd <- i
			<-even
		}
	}()

	// 偶数线程
	go func() {
		defer wg.Done()

		for i := 2; i <= 10; i += 2 {
			<-odd
			fmt.Println("even：", i)
			even <- i
		}
	}()

	wg.Wait()
}
