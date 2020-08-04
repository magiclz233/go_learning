package main

import "go_learning/go_base/ch19_reflect/reflect"

type UserFloat float64

func main() {
	var x UserFloat = 3.4
	reflect.Type(x)
	reflect.Value(x)
	//报错, 需要传入指针, 这样才能进行修改信息
	//reflect.UpdateValue(x)
	reflect.UpdateValue(&x)
	u := reflect.User{
		Id:   1,
		Name: "magic",
		Age:  23,
	}

	b := reflect.Boy{
		User: u,
		Addr: "西安",
	}
	reflect.StructReflect(u)

	reflect.StructAnon(b)
}
