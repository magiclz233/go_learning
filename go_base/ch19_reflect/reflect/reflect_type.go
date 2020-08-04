package reflect

import (
	"fmt"
	"reflect"
)

func Type(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("接口类型为:", t)

	// kind 可以获取具体类型
	k := t.Kind()
	fmt.Println("具体类型:", k)

	switch k {
	case reflect.Int:
		fmt.Printf("%v 类型为 %v\n", a, k)
	case reflect.Float64:
		fmt.Printf("%v 类型为 %v\n", a, k)
	case reflect.String:
		fmt.Printf("%v 类型为 %v\n", a, k)
	}

}
