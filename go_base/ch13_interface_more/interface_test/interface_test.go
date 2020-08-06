package interface_test

import (
	"fmt"
	"testing"
)

//自定义类型
type Code string

//接口
type Programmer interface {
	WriterHelloWorld(str string) Code
}

//接口实现类
type GoProgrammer struct {
}

func (p1 *GoProgrammer) WriterHelloWorld(str string) Code {
	//自定义类型也不支持自动转换
	return Code("GoProgrammer" + str)

}

//接口实现类
type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriterHelloWorld(str string) Code {
	return Code("JavaProgrammer" + str)
}

//工厂类 传入接口得不同实现类,实现不同的方法
func writeFirstProgram(p Programmer, str string) {
	fmt.Printf("%T %v\n", p, p.WriterHelloWorld(str))
}
func TestInterface(t *testing.T) {
	var a = "a"
	//必须是指针 不能是对象
	goTest := &GoProgrammer{}
	t.Log(goTest.WriterHelloWorld(a))
	javaTest := new(JavaProgrammer)
	t.Log(javaTest.WriterHelloWorld(a))

	writeFirstProgram(goTest, "go")
	writeFirstProgram(javaTest, "java")
}
