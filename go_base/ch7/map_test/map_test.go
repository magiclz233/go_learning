package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3}
	m1 := map[string]int{}
	m1["one"] = 1
	//m2 = make(map_ext[string]int,10)
	t.Log(m["one"])
	t.Log(m1["one"])
	//t.Log(m2["one"])
}

func TestAccessNotExistKey(t *testing.T) {
	m1 := map[int]int{}
	// 不存在 输出默认值0
	t.Log(m1[1])
	m1[2] = 0
	// 存在  输入key为2的value为0
	t.Log(m1[2])

	//两个返回值,v为value ok为是否存在
	if v, ok := m1[3]; ok {
		t.Log(v)
		t.Log("key 3 存在")
	} else {
		t.Log("key 3 不存在")
	}
}

//遍历map
func TestTravelMap(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3}
	//返回值为两个值  key value
	for k, v := range m {
		t.Log(k, v)
	}
}
