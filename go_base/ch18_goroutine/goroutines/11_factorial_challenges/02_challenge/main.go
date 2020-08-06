package main

import "fmt"

func main() {
	total := factorial(4)
	for v := range total {
		fmt.Println(v)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
