package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main")
	go longWait()
	go shortWait()
	fmt.Println("About the end of main()")
	time.Sleep(10 * 1e9)
	fmt.Println("ant the end of main()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(10 * 1e9) //sleep t seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) //sleep 2 seconds
	fmt.Println("End of shortWait()")
}
