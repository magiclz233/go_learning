package main

import "fmt"

func main() {
	c := incrementer()
	sum := puller(c)
	for v := range sum {
		fmt.Println(v)
	}
}

func incrementer() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for v := range c {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}
