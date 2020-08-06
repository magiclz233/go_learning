package map_ext

import "testing"

//map实现的工厂模式
func TestMapWithFunValue(t *testing.T) {
	//定义一个map key为int value为方法 传入参数 op int  输出int
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

//go没有实现set  可以用map实现set
func TestMapToSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	//set存储的是不能重复的值,go中map的key是不能重复的,创建一个key为需要保存的值,value为
	//布尔值的map,判断key是否存在,就能实现set的功能
	if mySet[1] {
		t.Log("key为 1 元素存在")
	} else {
		t.Log("key 为 1 不存在")
	}

	//输出map的数量
	t.Log(len(mySet))
	//删除map中指定的key
	delete(mySet, 1)
}
