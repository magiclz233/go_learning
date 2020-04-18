package type_test

import "testing"
// go的数据类型

//自定义别名类型
type MyInt int64

//go的类型转换
func TestImplicit(t *testing.T)  {
	var a int = 1
	var b int64
	//b = a 不支持隐式类型转换

	//只能显式的类型转换
	b = int64(a)

	//不支持相同类型的别名类型到原类型隐式转换  只能显式转换
	var c MyInt
	c = MyInt(b)

	t.Log(a,b,c)
}

//go的指针
func TestPoint(t *testing.T)  {
	a := 1
	aPtr := &a //0xc00005c290

	//错误  go不支持指针运算
	//aPtr = aPtr + 1
	
	t.Log(a,aPtr)
	t.Logf("%T %T",a,aPtr)

}

func TestString(t *testing.T)  {
	//string初始化为空字符串"",而不是null或者nil
	var s string
	t.Log("a"+s+"a")
	t.Log(len(s))

	//在对string进行非空判断时,要使用 string == "" 而不是string == nil
	if s == "" {
		t.Log("空字符串"+s)
	}
}
