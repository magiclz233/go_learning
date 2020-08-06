package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10000; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for v := range c {
			fmt.Println(v)
		}
		done <- true
	}()

	go func() {
		for v := range c {
			fmt.Println(v)
		}
		done <- true
	}()

	<-done
	<-done
}
