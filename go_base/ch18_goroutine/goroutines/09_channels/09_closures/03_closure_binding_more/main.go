package main

import "fmt"

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		v := v
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// 等待所有goroutine完成后再退出
	for _ = range values {
		<-done
	}
}

/*
Even easier is just to create a new variable,
using a declaration style that may seem odd but works fine in Go.

甚至更简单的是创建一个新变量，
使用看起来可能很奇怪但在Go中可以正常工作的声明样式。

SOURCE:
https://golang.org/doc/faq#closures_and_goroutines
*/
