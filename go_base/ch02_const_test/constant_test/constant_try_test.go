package constant_test

import "testing"

//声明常量
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Read = 1 << iota
	Writer
	Nothing
)

func TestConstDay(t *testing.T) {
	t.Log(Monday)
	t.Log(Tuesday)
	t.Log(Wednesday)

}

func TestConstTry(t *testing.T) {
	a := 7 //0111
	b := 1 //0001
	t.Log(Read)
	t.Log(Writer)
	t.Log(Nothing)
	t.Log(a ^ Writer)
	t.Log(b & Writer)
	t.Log(a&Read == Read, a&Writer == Writer, a&Nothing == Nothing)
	t.Log(b&Read == Read, b&Writer == Writer, b&Nothing == Nothing)

}
