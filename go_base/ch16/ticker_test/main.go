package main

import (
	"fmt"
	"time"
)

//Ticker 计时器  最小时间单位为ns(n纳秒, int64)
//Timer 定时器 只发送一次在传入的时间长度之后
func main() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
		default:
			fmt.Println(".....")
			time.Sleep(5e7)
		}
	}
}
