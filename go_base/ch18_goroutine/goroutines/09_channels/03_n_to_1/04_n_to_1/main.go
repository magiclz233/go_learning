package main

import "fmt"

func main() {

	for n := range test() {
		fmt.Println(n)
	}
}

func test() chan int {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < n; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()
	return c
}
