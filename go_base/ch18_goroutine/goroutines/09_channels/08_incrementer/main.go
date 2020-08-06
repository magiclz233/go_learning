package main

import "fmt"

func main() {
	c1 := incrementer("one: ")
	c2 := incrementer("two: ")
	p1 := puller(c1)
	p2 := puller(c2)
	fmt.Println("Final Count: ", <-p1+<-p2)
}

func incrementer(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
			fmt.Println(s, i)
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
