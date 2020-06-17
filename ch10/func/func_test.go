package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//多返回值函数
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

//一个计时函数,输入参数为函数,返回参数也为函数
//有点像spring的切面或者装饰者模式
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("函数运行时间:", time.Since(start).Seconds())
		return ret
	}
}

func TestFunc(t *testing.T) {

	a, b := returnMultiValues()
	t.Log(a, b)

	//多返回值如果想要省略其中的返回值,用 _ 代替
	c, _ := returnMultiValues()
	t.Log(c)
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
