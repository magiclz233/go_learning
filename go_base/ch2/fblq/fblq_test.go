package fblq

import (
	"fmt"
	"testing"
)

func TestFblqList(t *testing.T) {
	//var a int = 1
	//var b int = 1
	//var (
	//	a int = 1
	//	b = 1
	//)
	a := 1
	b := 1

	fmt.Println(a)
	t.Log(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()

}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	//交换数值
	//temp := a
	//a = b
	//b = temp
	//go中可以直接交换
	a, b = b, a
	t.Log(a, b)
}
