package panic_recover

import (
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	//defer func() {
	//	fmt.Println("finally")
	//}()

	//finally块相同作用
	defer func() {
		fmt.Println("defer")
		// recover_test 相当于java的try-catch,获取到所有的错误
		if err := recover(); err != nil {

			fmt.Println("recover_test from ", err)
		}
	}()
	fmt.Println("start")
	//直接退出
	//os.Exit(-1)

	//输入整个调用栈的信息
	//panic(errors.New("something wrong!"))

	//os.Exit后面的不会被执行 编译报错
	//fmt.Println("next")
}
