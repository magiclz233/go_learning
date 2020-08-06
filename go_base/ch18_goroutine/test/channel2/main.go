package main

import "fmt"

//协程的同步-关闭通道
func main() {
	ch := make(chan string)
	go sendData(ch)
	// getData不为协程 和main函数同一线程
	getData(ch)
}

func sendData(ch chan string) {
	ch <- "PHP"
	ch <- "Golang"
	ch <- "Java"
	ch <- "Python"
	close(ch)
}

func getData(ch chan string) {
	//如果加上sleep 则main函数直接执行完 go get协程还未执行 由于主线程已经关闭,所以不会输出
	//time.Sleep(2e9)
	//for {
	//	input, ok := <-ch
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("%s ", input)
	//}

	//使用 for-range 语句来读取通道是更好的办法，因为这会自动检测通道是否关闭
	for input := range ch {
		fmt.Printf("%s ", input)
	}
}
