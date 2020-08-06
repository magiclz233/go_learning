package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// 关闭channel
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

/*
如果不加 close(c)
该程序会报错 all goroutines are asleep - deadlock!
原因: 在循环插入chan完成之后,没有关闭通道 close(chan)
*/
