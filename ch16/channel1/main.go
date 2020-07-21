package main

import (
	"fmt"
	"time"
)

func main() {
	//ch := make(chan string)
	//go sendData(ch)
	//go getData(ch)
	//time.Sleep(1e9)

	//一个协程在无限循环中给通道发送整数数据。不过因为没有接收者，只输出了一个数字 0
	ch1 := make(chan int)
	go pump(ch1)
	go suck(ch1)
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "PHP"
	ch <- "Golang"
	ch <- "Java"
	ch <- "Python"
}

func getData(ch chan string) {
	var input string
	//如果加上sleep 则main函数直接执行完 go get协程还未执行 由于主线程已经关闭,所以不会输出
	//time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

//pump() 函数为通道提供数值，也被叫做生产者
func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

//为通道解除阻塞定义了 suck 函数来在无限循环中读取通道
func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
