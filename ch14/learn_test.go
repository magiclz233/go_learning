package ch14_test

import (
	"fmt"
	"os"
	"testing"
)

// 全局变量
var m = 100
var Max = 200

func TestVar(t *testing.T) {
	n := 100
	fmt.Printf("%d,%d", m, n)
}

func TestFor(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	for i := range arr {
		t.Log(i)
	}
}

func TestIf(t *testing.T) {

	i := []int{1, 3, 5, 7, 2, 4, 6, 8}

	for j := range i {
		if j%2 == 1 {
			fmt.Printf("奇数 %d;", j)
		}
		if j%2 == 0 {
			fmt.Printf("偶数 %d;", j)
		}
	}
}

func TestYunSuan(t *testing.T) {
	n := 2 & 3
	t.Log(n)
	i := 1 << 4
	t.Log(i)
	j := 2 >> 1
	t.Log(j)
	z := 2 ^ 3
	t.Log(z)
}

func Test_(t *testing.T) {
	buf := make([]byte, 1024)
	f, _ := os.Open("E:\\1.txt")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

func TestType(t *testing.T) {
	var a, b, c = 9223372036854775807, 2.1, "test"
	fmt.Printf("a type is %T\n", a)
	fmt.Printf("b type is %T\n", b)
	fmt.Printf("c type is %T\n", c)

}

func TestConst(t *testing.T) {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	area = LENGTH * WIDTH

	fmt.Printf("value of area: %d", area)
}

// 数组
func TestArr(t *testing.T) {
	var arr0 [5]int = [5]int{1, 2, 3}
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var arr2 = [...]int{1, 2, 3, 4}
	arr3 := [4]int{1, 2, 3, 4}
	var str = [5]string{2: "str1", 4: "str2"}
	a := [...]struct {
		name string
		age  int8
	}{
		{"user1", 18},
		{"user2", 20},
	}

	fmt.Println(arr0, arr1, arr2, arr3, str, a)
}

//多维数组
func TestArrMore(t *testing.T) {
	var arr0 [5][3]int
	var arr1 = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	a := [2][2]int{{1, 2}, {2, 3}}
	b := [...][3]int{{1, 2, 3}, {1, 2}} //第二维度不能用...
	fmt.Println(arr0, arr1, a, b)
}

//因为go是值拷贝,会造成性能问题  所以相对于数组的复用来说 slice(指针)或者数组指针用的更多

func arr(x [2]int) {
	fmt.Printf("x, %p\n", &x)
	x[1] = 1000
	fmt.Println(x) //[0 1000]
}

//go 数组值传递测试
func TestArrNum(t *testing.T) {
	a := [2]int{}
	fmt.Printf("a, %p\n", &a)
	arr(a)
	fmt.Println(a) // [0 0]
}

//多维数组遍历
func TestArrMoreIterator(t *testing.T) {
	f := [...][3]int{{1, 2, 3}, {4, 5, 6}, {14, 12, 11}}

	for k1, v1 := range f {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d) = %d ", k1, k2, v2)
		}
		fmt.Println()
	}
}

//数组拷贝和传参
func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

//传入的是arr1和arr2的指针,则printArr方法改动的就是该指针指向的arr1,arr2的数值
func TestArrCopy(t *testing.T) {
	arr1 := [5]int{}
	printArr(&arr1)
	fmt.Println(arr1)
	arr2 := [5]int{1, 2, 3, 4, 5}
	printArr(&arr2)
	fmt.Println(arr2)
}

//切片创建
func TestSlice1(t *testing.T) {
	var s1 []int
	s2 := []int{}
	s3 := make([]int, 2)
	fmt.Println(s1, s2, s3)
	s4 := make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	arr := [5]int{1, 2, 3, 4, 5}
	s6 := arr[1:4]
	fmt.Println(arr, s6)

	var arr1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var ss1 []int = arr1[:]
	var ss2 []int = arr1[2:]
	var ss3 = arr1[:5]
	var ss4 = arr1[2:5]

	t.Log("arr1: {},ss1: {},ss2: {},ss3: {},ss4: {}", arr1, ss1, ss2, ss3, ss4)

	var sl1 []int = make([]int, 10)
	var sl2 = make([]int, 10)
	sl3 := make([]int, 10, 10)

	fmt.Printf("make-sl1: %v\n", sl1)
	fmt.Printf("make-sl2: %v\n", sl2)
	fmt.Printf("make-sl3: %v\n", sl3)
}

func TestSlice2(t *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 5}
	s := data[2:4]

	s[0] += 100
	s[1] += 200

	fmt.Printf("data: %v\n", data)
	fmt.Printf("s: %v", s)

	s1 := []int{0, 1, 2, 3, 8: 100}
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 6, 8)
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 6)
	fmt.Println(s3, len(s3), cap(s3))
}

//使用指针操作切片
func TestSliceByPointer(t *testing.T) {
	s := []int{1, 2, 3, 4}

	//获取切片s下标2位置的数字的指针
	p := &s[2]

	//使用指针直接操作s[2]
	*p += 100
	fmt.Println(s)
}

//多维切片

func TestSliceTwo(t *testing.T) {
	data := [][]int{
		[]int{1, 2, 3},
		[]int{22, 23},
		[]int{11, 12, 13, 14},
	}

	fmt.Println(data)

	d := [5]struct {
		x int
	}{}

	s := d[:]
	d[1].x = 10
	s[2].x = 20

	fmt.Println(d)
	fmt.Printf("%p, %p\n", &d, &d[0])
}

//append追加slice
func TestSliceAppend(t *testing.T) {
	var a = []int{1, 2, 3}
	fmt.Printf("slice a : %p ,%v\n", a, a)
	b := []int{4, 5, 6}
	fmt.Printf("slice b :%p ,%v\n", b, b)

	c := append(a, b...)
	fmt.Printf("slice c :%p ,%v\n", c, c)

	d := append(c, 7)
	fmt.Printf("slice d :%p ,%v\n", d, d)

	e := append(d, 8, 9, 10)
	fmt.Printf("slice e :%p ,%v\n", e, e)

	s1 := make([]int, 0, 5)
	fmt.Printf("%p\n", s1)

	s2 := append(s1, 1)
	fmt.Printf("%p\n", s2)

	fmt.Println(s1, s2)
}

//超出原slice.cap的限制,就会重新分配新的地层数组,即使原数组未填满
func TestSliceCap(t *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	//[2:3] 表示[data[2],data[3]) [:2:3]表示 [:2] 然后cap最大为3
	s := data[:2:3]

	s = append(s, 100, 200)
	fmt.Println(s, data)
	//s append之后 超出了原来的数组范围  已经变成了新的数组,指针输出后不同
	fmt.Println(&s[0], &data[0])
}

//cap分配规律
func TestCap(t *testing.T) {
	s := make([]int, 0, 1)
	c := cap(s)

	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
		}
	}
}

//切片拷贝
func TestSliceCopy(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice s1: %v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("Slice s2: %v\n", s2)

	// s1的同位置copy到s2中
	copy(s2, s1)
	fmt.Printf("Copied Slice s1: %v\n", s1)
	fmt.Printf("Copied Slice s2: %v\n", s2)

	s3 := []int{1, 2, 3}
	fmt.Printf("slice s3: %v\n", s3)
	s3 = append(s3, s2...)
	fmt.Printf("appended slice s3: %v\n", s3)
	s3 = append(s3, 4, 5, 6)
	fmt.Printf("last slice s3: %v\n", s3)

}

//切片与地层数组copy之后的关系
func TestArrSliceCopy(t *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("data :", data)
	s1 := data[8:]
	s2 := data[:5]

	fmt.Println(s1, s2)

	copy(s2, s1)
	//因为切片其实是对底层数组的使用,上面copy,s1的数值改变了s2 所以data数组也发生了改变
	fmt.Println(data, s1, s2)
}

//切片遍历
func TestSliceIterator(t *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[:]
	slice[0] = 1
	for index, value := range slice {
		fmt.Printf("index: %v,index: %v\n", index, value)
	}

	var a = []int{1, 2, 3, 4, 5}
	fmt.Println(a, len(a))
	b := a[1:2] //[2]
	fmt.Println(b, len(b))
	//c虽然是b的切片 但是最终还是获取的a数组的值  c[:3] 表示以b的开头为开头取值
	c := b[0:3] // [2 3 4]
	fmt.Println(c, len(c))
}

// 切片和字符串
func TestSliceAndStr(t *testing.T) {
	str := "hello magic"
	s1 := str[:5]
	fmt.Println(s1)

	s2 := str[6:]
	fmt.Println(s2)

	str1 := "hello world"
	s := []byte(str1) //中文字符需要用[]rune(str1)
	s[6] = 'g'
	s = s[:7]
	s = append(s, 'o', '!')
	str1 = string(s)
	fmt.Println(str1)

	str2 := "你们的哥哥! magic"
	ss1 := []rune(str2)
	ss1[2] = '是'
	ss1[3] = '猪'

	ss1 = ss1[:4]
	str2 = string(ss1)
	fmt.Println(str2)

}

/*
golang slice data[:6:8] 两个冒号的理解

常规slice , data[6:8]，从第6位到第8位（返回6， 7），长度len为2， 最大可扩充长度cap为4（6-9）

另一种写法： data[:6:8] 每个数字前都有个冒号， slice内容为data从0到第6位，长度len为6，最大扩充项cap设置为8

a[x:y:z] 切片内容 [x:y] 切片长度: y-x 切片容量:z-x
*/
func TestSome(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s1 := slice[4:7]
	s2 := slice[:4:6]
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
}

func TestMap1(t *testing.T) {
	ages := map[string]int{}
	ages["magic"] = 1
	ages["root"] = 2
	// ok用来判断你获取的元素是否存在 不存在返回false
	age, ok := ages["r"]
	fmt.Println(age, ok)

	test := make(map[string]int)
	test["magic"] = 1
	test["root"] = 2

	isOk := equals(ages, test)
	fmt.Println(isOk)
}

func equals(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			fmt.Println(!ok)
			return false
		}
	}
	return true
}

//函数
func TestFunc1(t *testing.T) {
	nextInt := intSeq()
	//不同的创建实例之后,数据互不相同

	fmt.Println(nextInt()) //1
	fmt.Println(nextInt()) //2
	fmt.Println(nextInt()) //3

	newInt := intSeq()
	fmt.Println(newInt()) //1
}

//闭包,匿名函数
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

//可变参数加多返回值函数
func TestFunc2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	total, multiply := sumAndMultiply(nums...)
	fmt.Println(total, multiply)

	fact := fact(7)
	fmt.Println(fact)
}

//可变参数加多返回值函数
func sumAndMultiply(nums ...int) (int, int) {
	fmt.Print(nums, " ")
	total := 0
	multiply := 1
	for _, num := range nums {
		total += num
	}
	for _, num := range nums {
		multiply *= num
	}
	return total, multiply
}

// 递归
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

//指针
func TestPointer(t *testing.T) {
	i := 1
	fmt.Println("初始化的i : ", i) // 1

	val(i)
	fmt.Println("值传递之后,main函数中的 i : ", i) // 1

	point(&i)
	fmt.Println("引用传递之后,main函数中的 i : ", i) // 0

}

func val(val int) {
	val = 0
}

func point(point *int) {
	*point = 0
}
