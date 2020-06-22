package pointer_test

import (
	"fmt"
	"testing"
)

// 指针

func TestPointer1(t *testing.T) {
	a := 10
	b := &a

	fmt.Printf("a: %d pointer: %p \n", a, &a)
	fmt.Printf("b %p, type: %T\n", b, b)

	// 指针取值 *b
	c := *b
	fmt.Println(a, b, c)             //10 0xc00005c2c0 10
	fmt.Printf("type of c:%T\n", c)  // int
	fmt.Printf("value of c:%v\n", c) // 10
}

//值传递和引用传递

func modify1(x int) int {
	x += 100
	fmt.Println("modify1方法内部x: ", x)
	return x
}

func modify2(x *int) int {
	*x += 100
	fmt.Println("modify2方法内部x: ", *x)
	return *x
}

func TestPointer2(t *testing.T) {
	x := 1
	fmt.Println("初始化x: ", x)
	modify1(x)
	fmt.Println("test modify1 x: ", x)
	modify2(&x)
	fmt.Println("test modify2 x: ", x)
}

// 空指针

func TestNilPointer(t *testing.T) {
	var p *string
	fmt.Println("p:", p)

	////报错空指针
	//fmt.Printf("p的值为 %s\n", *p)

	magic := "magic"

	p = &magic

	fmt.Println("*p :", *p)

	if p != nil {
		fmt.Println("非空指针p")
	} else {
		fmt.Println("空指针")
	}
}

//new
func TestPointer3(t *testing.T) {
	//没有分配内存直接初始化的指针  报错panic
	//var a *int
	//*a = 100
	//fmt.Println(*a)
	//
	//var b map[string]int
	//b["测试"] = 100
	//fmt.Println(b)

	// func new(Type) *Type
	//new方法接收type 返回初始化好的一个指针

	var s *int
	s = new(int)
	*s = 10
	fmt.Println(*s)

	a := new(int)
	b := new(string)
	c := new(bool)
	d := new(map[string]int)
	e := new([]int)
	arr := [...]int{1, 2, 3, 4}
	f := &arr
	g := new([5]int)

	fmt.Println(*f)

	fmt.Printf("%T,%T,%T,%T,%T,%T\n", a, b, c, d, e, g)
	fmt.Println(*a, *b, *c, *d, *e, *g)
}

// make: func make(t Type, ...IntegerType) Type
func TestMakePointer(t *testing.T) {
	var b map[string]int
	b = make(map[string]int, 1)
	b["测试"] = 100
	b["test"] = 200
	fmt.Println(b)
	fmt.Println(len(b))
}

func pointerTest(x *int) *int {
	*x = 100
	fmt.Println(*x)
	return x
}

func TestPointer4(t *testing.T) {
	a := new(int)
	a = pointerTest(a)
	fmt.Println(*a)
	b := *a
	fmt.Println(b)

}

func TestMapSet(t *testing.T) {
	mySet := map[string]bool{}
	mySet["magic"] = true
	mySet["root"] = true
	var a = "magic"
	for k, v := range mySet {
		if k == a {
			fmt.Println(v)
		}
	}
}

//map 类型的切片

func TestMapSlice(t *testing.T) {
	mapSlice := make([]map[string]string, 3)
	for i, v := range mapSlice {
		fmt.Printf("index: %d, value: %v\n", i, v)
	}

	fmt.Println("初始化")

	mapSlice[0] = make(map[string]string, 3)
	mapSlice[0]["name"] = "张三"
	mapSlice[0]["sex"] = "男"
	mapSlice[0]["age"] = "23"

	for i, v := range mapSlice {
		fmt.Printf("index: %d, value: %v\n", i, v)
	}
}

//值为切片的map
func TestSliceMap(t *testing.T) {
	sliceMap := make(map[string][]string, 10)
	fmt.Println(sliceMap)

	key := "中国"

	value, ok := sliceMap[key]

	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "西安")

	sliceMap[key] = value
	fmt.Println(sliceMap)
}
