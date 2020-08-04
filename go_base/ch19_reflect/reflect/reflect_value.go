package reflect

import (
	"fmt"
	"reflect"
)

func Value(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Printf("a 的值为: %v\n", a)

	k := v.Kind()

	fmt.Printf("a 的具体类型为: %v\n", k)

	switch k {
	case reflect.Int:
		fmt.Printf("%v 值为 %v,类型为 %v\n", a, v, k)
	case reflect.Float64:
		fmt.Printf("%v 值为 %v,类型为 %v\n", a, v, k)
	case reflect.String:
		fmt.Printf("%v 值为 %v,类型为 %v\n", a, v, k)
	}
}
