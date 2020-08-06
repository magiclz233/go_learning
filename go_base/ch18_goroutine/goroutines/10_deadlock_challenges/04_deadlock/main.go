package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		fmt.Println(<-c)
	}
}

/*
这样的程序会报错 all goroutines are asleep - deadlock!
两个缺点:
	1: 在chan添加完成所有值后,没有关闭通道 close(chan)
	2: for{} 为死循环, 会循环遍历 <- chan的第一个值 该程序为循环遍历 0 不停止
*/
