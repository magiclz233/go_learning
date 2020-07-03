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
