package fn

import (
	"fmt"
	"testing"
)

//函数可变长参数
//不确定多少个传入的参数相加
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

//函数执行函数 defer块
//类似于try-catch-finally中的finally块

func Clear() {
	fmt.Println("关闭资源或者释放锁")
}
func TestDefer(t *testing.T) {
	//虽然在上面写着，但是会先执行下面得fmt方法
	defer Clear()        // 执行顺序：2
	t.Log("log start")   //3
	fmt.Println("start") //1
	//报错，后面的无法继续执行 但是 defer 还是可以的
	panic("error")
	//panic后的函数不可达 无法调用
	//fmt.Println("next start")
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5))
	t.Log(Sum(1, 2, 3, 4))
}
