package main

import "go_learning/go_base/ch19_reflect/reflect"

type UserFloat float64

func main() {
	var x UserFloat = 3.4
	reflect.Type(x)
	reflect.Value(x)
}
