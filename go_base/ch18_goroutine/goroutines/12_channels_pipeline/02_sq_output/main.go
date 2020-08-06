package main

import "fmt"

func main() {
	for n := range sq(sq(gen(2, 3, 4))) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int, len(nums))
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
