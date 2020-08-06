package method_test

import (
	"fmt"
	"math"
	"testing"
	"time"
)

//方法声明

type Point struct {
	X, Y float64
}

// 通用函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Point结构体的方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func TestDistance(t *testing.T) {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // 5 function call
	fmt.Println(p.Distance(q))  // 5 method call
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i, _ := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func TestPath(t *testing.T) {
	p := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(p.Distance())
}

// 基于指针对象的方法

//按照比例方法缩小
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func TestPointerMethod(t *testing.T) {
	r := &Point{1, 2}

	//非指针方法 不会改变 r值
	q := Point{4, 6}
	fmt.Println(r.Distance(q))

	r.ScaleBy(2)
	fmt.Println(*r)

	p := Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p)
	// 方法
	fmt.Println(pptr.Distance(q))
	// 指针
	fmt.Println((*pptr).Distance(q))

	p1 := Point{1, 2}
	(&p1).ScaleBy(2)
	fmt.Println(p1)
}

/*
1. 不管你的method的receiver是指针类型还是非指针类型，都是可以通过指针/非指针类型进行调用的，编译器会帮你做类型转换。
2. 在声明一个method的receiver该是指针还是非指针类型时，你需要考虑两方面的因素，第一方面是这个对象本身是不是特别大，
如果声明为非指针变量时，调用会产生一次拷贝；第二方面是如果你用指针类型作为receiver，那么你一定要注意，这种指针类型
指向的始终是一块内存地址，就算你对其进行了拷贝。熟悉C或者C++的人这里应该很快能明白。
*/

// Nil也是一个合法的接收器类型

type IntList struct {
	Value int
	Tail  *IntList
}

func (list IntList) Sum() int {
	return list.Value + list.Tail.Sum()
}

//闭包 相当于匿名函数
func returnNum() func(a, b int) (int, int) {
	return func(a, b int) (int, int) {
		return a, b
	}
}

func TestSFunc(t *testing.T) {
	numFunc := returnNum()
	a, b := numFunc(1, 2)
	fmt.Println(a, b)
}

// 通过嵌入结构体来扩展类型

func TestSelect(t *testing.T) {
	tick := time.Tick(time.Millisecond * 100)
	boom := time.After(time.Millisecond * 500)
	for {
		select {
		// 100毫秒运行一次
		case <-tick:
			fmt.Println("tick.")
		//	500毫秒运行榆次
		case <-boom:
			fmt.Println("boom.")
		//	50毫秒运行一次
		default:
			fmt.Println(".......")
			time.Sleep(time.Millisecond * 50)
		}
	}
}
