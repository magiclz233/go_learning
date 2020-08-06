package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var count int
var mutex sync.Mutex

func main() {
	wg.Add(3)
	go incrementer("One:")
	go incrementer("Two:")
	go incrementer("Three:")
	wg.Wait()
	fmt.Println("Final count: ", count)
}

func incrementer(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		count++
		fmt.Println(s, i, "Count:", count)
		mutex.Unlock()
	}
	wg.Done()
}
