package interface_test

import (
	"fmt"
	"testing"
)

// 实现多个接口
type Say interface {
	say()
}

type Mover interface {
	move()
}

type dog1 struct {
	Name string
}

func (d *dog1) say() {
	fmt.Println("dog1 say")
}

func (d *dog1) move() {
	fmt.Println("dog1 move")
}

func (d dog) move() {
	fmt.Println("dog move")
}

func (c cat) move() {
	fmt.Println("cat move")
}

func TestOneByTwoInterface(t *testing.T) {
	var x Say
	var y Mover
	var d = dog1{Name: "dog1"}
	x = &d
	y = &d
	x.say()
	y.move()

	// 多个类型实现同一接口
	var m Mover

	var c = cat{}
	var dogs = dog{}
	m = c
	m.move()
	m = dogs
	m.move()

}

// 接口嵌套

type Sayer1 interface {
	say()
}

type Mover1 interface {
	move()
}

//嵌套的接口
type Animal interface {
	Sayer1
	Mover1
}

type Pig struct {
	Name string
}

type Rabbit struct {
	Name string
}

func (p Pig) say() {
	fmt.Println("Pig say")
}
func (p Pig) move() {
	fmt.Println("Pig move")
}

func (p Rabbit) say() {
	fmt.Println("Rabbit say")
}

func (p Rabbit) move() {
	fmt.Println("Rabbit move")
}

func TestInterfaceNest(t *testing.T) {
	var pig = Pig{Name: "佩奇"}
	var rabbit = Rabbit{Name: "小白"}

	var a Animal
	a = pig
	a.move()
	a.say()
	a = rabbit
	a.say()
	a.move()
}

// 空接口

// 空接口可以存储任意类型值, 相当于java中的Object 用空接口当参数可以接收任何类型的数据
func show(a interface{}) {
	fmt.Printf("type: %T, value: %t\n", a, a)
}

func TestNullInterface(t *testing.T) {
	show("type string")
	show(123)
	show(111.11)
	show([]int{1, 2, 3})

	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "magic"
	studentInfo["age"] = 18
	studentInfo["sex"] = "男"
	studentInfo["married"] = false

	fmt.Println(studentInfo)
}

// 类型断言

// 该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。

func switchType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string, value is %v\n", v)
	case int:
		fmt.Printf("x is a int, value is %d\n", v)
	case bool:
		fmt.Printf("x is a bool, value is %v\n", v)
	case float64:
		fmt.Printf("x is a float64, value is %v\n", v)
	default:
		fmt.Println("un support type!")
	}
}

func TestInterface(t *testing.T) {
	var x interface{}
	x = "magic"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型错误")
	}

	x = 123
	i, ok1 := x.(int)
	if ok1 {
		fmt.Println(i)
	} else {
		fmt.Println("类型错误")
	}

	switchType("magic")
	switchType(123)
	switchType(true)
	switchType(11.111)
	switchType([...]int{1, 2, 3})
}
