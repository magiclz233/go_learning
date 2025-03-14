package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementer("Foo:")
	go incrementer("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementer(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

/*
	counter 未加锁 同时启动两个协程访问 incrementer方法
	counter 因为未加锁的原因会出现当一个正在执行++操作时,另外一个获取到了之前的值,
	最后导致counter最终结果不为预期的40
	如何解决: 1. sync.Mutex 互斥锁
			2. atomic 原子类
*/
// go run -race struct_two_test.go
// vs
// go run struct_two_test.go
