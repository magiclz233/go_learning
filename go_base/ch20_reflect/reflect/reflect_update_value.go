package reflect

import (
	"fmt"
	"reflect"
)

func UpdateValue(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Float64:
		v.SetFloat(6.9)
	case reflect.String:
		v.SetString("magic")
	case reflect.Ptr:
		// 指针类型
		// Elem()获取地址指向的值
		v.Elem().SetFloat(7.7)
		fmt.Println("case:", v.Elem().Float())
		// 地址
		fmt.Println(v.Pointer())
	}
}
