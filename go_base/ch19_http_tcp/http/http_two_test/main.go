package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"https://www.bookstack.cn/",
	"http://www.baidu.com",
	"https://juejin.im/",
}

func main() {
	for _, url := range urls {
		res, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, " : ", res.Status, res.StatusCode)
	}
}
