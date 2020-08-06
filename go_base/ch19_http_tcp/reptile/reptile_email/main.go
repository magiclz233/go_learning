package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	// \d是数字
	reQQEmail = `(\d+)@qq.com`
)

func GetEmail() {
	res, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleError(err, "http.GET url")
	defer res.Body.Close()

	// 读取页面数据
	pageBytes, err := ioutil.ReadAll(res.Body)
	HandleError(err, "ioutil.ReadAll")
	// 字节转字符串
	pageStr := string(pageBytes)
	// 过滤数据 获取QQ邮箱
	re := regexp.MustCompile(pageStr)
	// -1代表取全部
	results := re.FindAllStringSubmatch(pageStr, -1)

	for _, result := range results {
		fmt.Println("email: ", result[0])
		fmt.Println("qq: ", result[1])
	}
}

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
func main() {
	GetEmail()
}
