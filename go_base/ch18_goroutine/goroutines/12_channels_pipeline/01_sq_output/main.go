package main

import "fmt"

func main() {
	// 设置管道pipeline
	c := gen(2, 3, 4)
	out := sq(c)

	//顺序输出
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
	fmt.Println(<-out)
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
