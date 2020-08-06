package array_test

import "testing"

//声明数组
func TestArray1(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [2][2]int{{1, 2}, {3, 4}}
	arr3 := [...]int{1, 2, 3}
	arr4 := [...][]int{{1, 2}, {3, 4}}
	t.Log(arr, arr1, arr2, arr3, arr4)
}

//数组的循环展示
func TestArray2(t *testing.T) {
	arr1 := [...]int{1, 2, 3, 4, 5}
	//普通for循环
	for i := 0; i < len(arr1); i++ {
		t.Log(arr1[i])
	}
	//相当于foreach,idx为索引值,e为元素值
	for idx, e := range arr1 {
		t.Log(idx, e)
	}
	//下划线表示不想使用这个结果,但是必须进行占位
	for _, e := range arr1 {
		t.Log(e)
	}
}

//数组的截取
func TestArray3(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 7}

	//截取值时为左闭右开,第一个值为起始索引(包含),第二个值为结束索引(不包含)

	t.Log(arr[1:2])
	t.Log(arr[1:3])
	t.Log(arr[1:len(arr)])
	t.Log(arr[1:])
	t.Log(arr[:3])
	t.Log(arr[:])
	//不支持负数
	//t.Log(arr[:-1])
}
