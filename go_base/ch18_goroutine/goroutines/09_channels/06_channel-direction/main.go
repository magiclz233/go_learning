package main

import "fmt"

func main() {
	c := incrementer()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementer() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

/*

chan作为函数返回值的方式有3种:（chan int）、（<- chan int）、（chan <- int），
分别代表（可读可写的管道）、（只读管道）、（只写管道），只读管道不能close()，只写管道可以close()

The optional <- operator specifies the channel direction, send or receive.
If no direction is given, the channel is bidirectional.
https://golang.org/ref/spec#Channel_types
*/
