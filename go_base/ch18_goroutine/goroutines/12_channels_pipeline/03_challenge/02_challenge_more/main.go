package main

import "fmt"

func main() {
	for n := range factorialGo(gen(20)) {
		fmt.Println(n)
	}
}

func gen(n int) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorialGo(in chan int) chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			total := 1
			for i := v; i > 0; i-- {
				total *= i
			}
			out <- total
		}
		close(out)
	}()

	return out
}

/*
	使用管道并行运行20个阶乘
*/
