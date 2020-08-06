package interface_test

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

//一个计时函数,输入参数为函数,返回参数也为函数
//有点像spring的切面或者装饰者模式
//func timeSpent(inner func(op int) int) func(op int) int {
//	return func(n int) int {
//		start := time.Now()
//		ret := inner(n)
//		fmt.Println("函数运行时间:",time.Since(start).Seconds())
//		return ret
//	}
//}
//IntConv代替 func(op int) int
func timeSpent(inner func(op int) int) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("函数运行时间:", time.Since(start).Seconds())
		return ret
	}
}

//测试的时间函数
func slowFunc(op int) int {
	fmt.Println(op)
	time.Sleep(time.Second * 2)
	return op
}

//测试函数为参数和返回值的方法
func TestFunc2(t *testing.T) {

	spent := timeSpent(slowFunc)
	t.Log(spent(10))
}
