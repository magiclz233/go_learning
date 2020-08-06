package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	// FAN OUT
	// 从同一个chan 往多个sq work传入
	c1 := sq(in)
	c2 := sq(in)

	// FAN IN
	// 消耗来自多个通道的合并输出
	for v := range merge(c1, c2) {
		fmt.Println(v)
	}
}

func gen(nums ...int) chan int {
	fmt.Printf("Type Of nums: %T\n", nums)

	out := make(chan int)
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
		for v := range in {
			fmt.Println(v)
			out <- v * v
		}
		close(out)
	}()
	return out
}

// 将传入的所有通道的信息全部载入到chan out 然后返回
func merge(cs ...chan int) chan int {
	fmt.Printf("Type Of cs: %T\n", cs)

	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for v := range ch {
				out <- v
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
