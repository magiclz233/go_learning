package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
	fmt.Println(<-c)
}

/*
go func 给chan里面循环输入了0-9
fmt.Println(<- c)  <- c 只获取到了第一个输入的  要输出全部chan 用for-range
*/
