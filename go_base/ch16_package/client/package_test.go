package client

import (
	"fmt"
	"go_base/ch16_package/series"
	"testing"
)

func TestPackage(t *testing.T) {
	fibList := make([]int, 0)
	var err error
	// 访问series包中的函数  首字符大写的可以包外访问
	fibList, err = series.GetFibonacci(20)
	if err != nil {
		fmt.Println("fail", err)
		return
	}
	fmt.Printf("fibList: %v\n", fibList)

	// 小写的不能被包外访问
	//series.square(5)
	t.Log(series.Square(5)) // 25
}
