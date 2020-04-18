package condition

import "testing"

func TestIfMultiSec(t *testing.T) {
	/*
		前面 a := 1 == 1 是对a的声明加判断 后面 a 为返回值为true挥着false
		然后来进行if判断语句使用
	 */
	if a := 1 == 1;a{
		t.Log("1==1")
	}
}

/*
	someFunc() 方法是一个多返回值方法,第一个是返回结果,第二个是是否有错误

	if进行判断 如果err不为nil,则进入if执行程序
	如果为nil 则进入else
 */

/*
func TestManyReturn(t *testing.T) {

	if v,err := someFunc(); err{
		...
	}else {
		...
	}
}
 */
