package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

func HandlerError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func DownloadFile(url string, fileName string) (ok bool) {
	res, err := http.Get(url)
	baseSrc := `G:\picture\`
	HandlerError(err, "http.get.url")
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	HandlerError(err, "res.Body")
	fileName = baseSrc + fileName
	err = ioutil.WriteFile(fileName, bytes, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}

/*
 	并发爬虫思路
	1: 初始化数据管道
	2: 爬虫写出: 26个协程向管道添加图片链接
	3: 任务统计协程: 检查26个任务是否都完成, 完成则关闭数据管道
	4: 下载协程,从管道中读取链接并下载
*/

var (
	chanImageUrls chan string
	waitGroup     sync.WaitGroup
	chanTask      chan string
	reImg         = `https://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func main() {
	// 1. 初始化管道
	chanImageUrls = make(chan string, 1000000)
	chanTask = make(chan string, 26)
	// 2. 爬虫协程
	for i := 1; i < 27; i++ {
		waitGroup.Add(1)
		go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}
	// 3. 任务统计协程 统计26个协程是否都任务完成, 完成则关闭管道
	waitGroup.Add(1)
	go CheckOK()
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go DownloadImg()
	}
	waitGroup.Wait()
}

func DownloadImg() {
	for url := range chanImageUrls {
		fileName := GetFileNameFromUrl(url)
		ok := DownloadFile(url, fileName)
		if ok {
			fmt.Printf("%s 下载成功!!!!\n", fileName)
		} else {
			fmt.Printf("%s 下载失败!\n", fileName)
		}
	}
	waitGroup.Done()
}

// 截取url名字
func GetFileNameFromUrl(url string) string {
	// 返回最后一个 / 位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	filename := url[lastIndex+1:]
	timePreFix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePreFix + "_" + filename
	return filename
}

// 任务统计协程
func CheckOK() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 爬取成功\n", url)
		count++
		if count == 26 {
			close(chanImageUrls)
			break
		}
	}
	waitGroup.Done()
}

// 爬图片链接到管道
// url传入整页链接
func getImgUrls(url string) {
	urls := getImages(url)
	for _, url := range urls {
		chanImageUrls <- url
	}
	// 标识当前协程完成
	// 每完成一个任务写一条数据
	//用于监控知道已经完成了多少任务
	chanTask <- url
	waitGroup.Done()
}

// 获取当前页图片链接
func getImages(url string) (urls []string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("找到的结果数量: %v\n, 分别为: %v\n", len(results), results)
	for _, result := range results {
		url = result[0]
		urls = append(urls, url)
	}
	return urls
}

// 根据url获取内容
func GetPageStr(url string) (pageStr string) {
	res, err := http.Get(url)
	HandlerError(err, "http.GET url")
	defer res.Body.Close()
	//读取页面数据
	pageBytes, err := ioutil.ReadAll(res.Body)
	HandlerError(err, "ioutil.ReadAll")
	// 字节转字符串
	pageStr = string(pageBytes)
	return
}
