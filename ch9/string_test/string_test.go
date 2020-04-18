package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	//初始化为 "" 空字符串
	t.Log(s)
	s = "hello"
	t.Log(len(s))

	//s[1] = '3' //编译错误,string是不可变的byte数组
	s = "\xE4\xB8\xA5"

	//一个汉字的string长度为3
	t.Log(s,len(s))

	s = "中"
	// len函数返回的是byte数
	t.Log(len(s))

	c := []rune(s)
	t.Log(len(c))

	t.Logf("中 unicode %x",c[0])
	t.Logf("中 UTF8 %x",s)
}

func TestStringFor(t *testing.T) {

	s := "我是你们的哥哥"
	for _,c := range s{
		t.Logf("%[1]c %[1]x",c)
	}
}
