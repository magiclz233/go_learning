package reflect

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u *User) Hello() {
	fmt.Printf("Hello %s", u.Name)
}

func StructReflect(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型: ", t)
	fmt.Println("Name字符串类型: ", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("值: ", v)

	// 获取所有属性
	// 获取结构体字段个数 t.NumField()

	for i := 0; i < t.NumField(); i++ {
		// 获取每个字段
		f := t.Field(i)
		fmt.Printf("%s : %v", f.Name, f.Type)
		// 获取字段的值信息
		// Interface() 获取字段对应的值
		val := v.Field(i).Interface()

		fmt.Println(" : ", val)
	}

	fmt.Println("================方法====================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}
