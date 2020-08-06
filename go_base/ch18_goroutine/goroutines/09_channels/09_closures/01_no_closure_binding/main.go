package main

import "fmt"

func main() {
	done := make(chan bool)
	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// 从 chan 取出值, 等待所有的协程完成
	for _ = range values {
		<-done
	}
}

/*
	for-range values 当你在循环时, 里面的协程 go func(){}() 未启动,循环完成时启动起来
	所以三个协程输出的v都是c
*/
