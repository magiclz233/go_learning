package reflect

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int    `json:"id1" db:"id2"`
	Name string `json:"name1" db:"name2"`
	Age  int    `json:"age1" db:"age2"`
}

func (u *User) Hello(name string) {
	fmt.Printf("Hello %s, %s\n", name, u.Name)
}

type Boy struct {
	User
	Addr string
}

// 获取结构体匿名字段
func StructAnon(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型: ", t)

	fmt.Printf("%#v\n", t.Field(0))

	fmt.Printf("%#v\n", reflect.ValueOf(o).Field(0))
}

// 获取结构体的类型和值
func StructReflect(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型: ", t)
	fmt.Println("字符串类型: ", t.Name())

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

// 修改结构体的值, 需要传入结构体的指针
func StructUpdate(o interface{}, name string) {
	v := reflect.ValueOf(o)
	//获取指针指向的元素
	v = v.Elem()
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String {
		f.SetString(name)
	}
}

func StructFunc(o interface{}, method string) {
	v := reflect.ValueOf(o)
	// 获取方法
	m := v.MethodByName(method)
	// 构建一些参数
	args := []reflect.Value{reflect.ValueOf("magic")}
	// 没参数的情况下: var args2 []reflect.Value
	// 调用方法需要传入方法的参数
	m.Call(args)
}

func StructTag(o interface{}, tag ...string) {
	v := reflect.ValueOf(o)
	t := v.Type()
	for i := 0; i < t.Elem().NumField(); i++ {
		f := t.Elem().Field(i)
		for _, s := range tag {
			fmt.Printf("%v 的 tag %v 是 %v\n", f.Name, s, f.Tag.Get(s))
		}
	}
}
