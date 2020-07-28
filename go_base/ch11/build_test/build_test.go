package build_test

import (
	"fmt"
	"testing"
	"unsafe"
)

//实体类的生成
type Employee struct {
	Id   string
	Name string
	Age  int
}

//通过实例访问 会复制Employee到内存新的地方,不建议使用,大量使用会造成内存浪费
func (e Employee) String() string {
	fmt.Printf("存储地址1: %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id:%s,Name:%s,Age:%d", e.Id, e.Name, e.Age)
}

//通过指针访问 使用指针直接指向生成的Employee对象原地址 无内存浪费
func (e *Employee) String1() string {
	fmt.Printf("存储地址2: %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id:%s,Name:%s,Age:%d", e.Id, e.Name, e.Age)
}

//测试实例访问和指针访问
func TestEmployeeFunc(t *testing.T) {
	e := Employee{"1", "MAGIC", 23}
	fmt.Printf("存储地址: %x", unsafe.Pointer(&e.Name))

	t.Logf(e.String())
	t.Logf(e.String1())
}

//对象的创建
func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"1", "magic", 23}
	e1 := Employee{Name: "ruby", Age: 23}
	//获取的是Employee的指针
	e2 := new(Employee)
	e2.Id = "2"
	e2.Name = "skt"
	e2.Age = 23
	t.Log(e)
	t.Log(e1)
	t.Log(e2)
	t.Log(e1.Id)
	// %T 返回当前位置参数的类型
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)

}
