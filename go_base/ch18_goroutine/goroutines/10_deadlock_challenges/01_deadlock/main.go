package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

/*
 chan 必须在协程中使用, 主线程中使用会报错 all goroutines are asleep - deadlock!
*/
