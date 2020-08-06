package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var count int64

func main() {
	wg.Add(2)
	go incrementer("One:")
	go incrementer("Two:")
	wg.Wait()
	fmt.Println("Final Count: ", count)
}

func incrementer(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		j := atomic.AddInt64(&count, 2)
		fmt.Println(j)
		fmt.Println(s, i, "Count : ", atomic.LoadInt64(&count))
	}
	wg.Done()
}
