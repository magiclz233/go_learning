package interface_test

import (
	"fmt"
	"testing"
)

var i = 5
var str = "ABC"

// 空接口
type Any interface {
}

type Person struct {
	name string
	age  int
}

func TestNullInterface(t *testing.T) {
	var val Any
	val = 5
	fmt.Printf("val: %v\n", val)
	val = str
	fmt.Printf("val: %v\n", val)
	per1 := new(Person)
	per1.name = "magic"
	per1.age = 23
	val = per1
	fmt.Printf("val: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}

type specialString string

var whatIsThis specialString = "hello"

func TypeSwitch() {
	testFunc := func(any interface{}) {
		switch v := any.(type) {
		case bool:
			fmt.Printf("any %v is a bool type", v)
		case int:
			fmt.Printf("any %v is an int type", v)
		case float32:
			fmt.Printf("any %v is a float32 type", v)
		case string:
			fmt.Printf("any %v is a string type", v)
		case specialString:
			fmt.Printf("any %v is a special String!", v)
		default:
			fmt.Println("unknown type!")
		}
	}
	testFunc(whatIsThis)
}

func TestSpecialString(t *testing.T) {
	TypeSwitch()
}

type Element interface{}

type Vector struct {
	a []Element
}

func (p *Vector) Get(i int) Element {
	return p.a[i]
}

func (p *Vector) Set(i int, e Element) {
	p.a[i] = e
}

func (p *Vector) Type(i int) {
	m := p.a[i]
	switch v := m.(type) {
	case bool:
		fmt.Printf("bool, %v, %T", v)
	}
}

func TestElement(t *testing.T) {
	el := new(Vector)
	el.Set(0, "magic")
	el.Set(1, 1)
	el.Set(2, [...]int{1, 2, 3})
	el.Set(3, true)

}
