package main

import "fmt"

func main() {
	done := make(chan bool)
	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func(u string) {
			fmt.Println(u)
			done <- true
		}(v)
	}

	// 从 chan 取出值, 等待所有的协程完成
	for _ = range values {
		<-done
	}
}

/*
go func(){}()改为 go func(u string){}(v) fmt.Println(u) 每次循环都会传进去
保证可以获取到

要将v的当前值绑定到每个闭包启动时，
必须在每次迭代中修改内部循环以创建一个新变量。
一种方法是将变量作为参数传递给闭包。

在此示例中，v的值作为参数传递给匿名函数。
然后可以在函数内部将该值作为变量u访问。

To bind the current value of v to each closure as it is launched,
one must modify the inner loop to create a new variable each iteration.
One way is to pass the variable as an argument to the closure.

In this example, the value of v is passed as an argument to the anonymous function.
That value is then accessible inside the function as the variable u.

SOURCE:
https://golang.org/doc/faq#closures_and_goroutines
*/
