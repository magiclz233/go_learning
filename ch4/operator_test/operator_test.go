package operator_test

import "testing"

const (
	Read = 1 << iota
	Writer
	Nothing
)

//go为值传递,长度相同的可以进行比较
func TestOperator(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1,2,3,4,5}
	d := [...]int{1, 2, 3, 4}

	//长度不同直接报错
	//t.Log(a == c)
	t.Log(a == b) //false
	t.Log(a == d) //true
	t.Log(a[1] << 2)
}

func TestBitClear(t *testing.T)  {
	a := 7 //0111
	a = a &^ Read
	a = a &^ Writer
	a = a &^ Nothing
	b := 1 //0001
	t.Log(Read)
	t.Log(Writer)
	t.Log(Nothing)
	t.Log(a ^ Writer)
	t.Log(b & Writer)
	t.Log(a&Read == Read,a&Writer == Writer, a&Nothing == Nothing)
	t.Log(b&Read == Read,b&Writer == Writer, b&Nothing == Nothing)

}