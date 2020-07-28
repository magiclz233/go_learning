package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	//可变长度的数组
	var s0 []int

	//len()获取切片中可访问的元素的个数,cap()获取切片的容量
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2]) // 0 0 0
	//报错,后面两个无法访问
	//t.Log(s2[0],s2[1],s2[2],s2[3],s2[4])
	s2 = append(s2, 1)
	s2 = append(s2, 2)
	t.Log(s2[0], s2[1], s2[2], s2[3], s2[4]) // 0 0 0 1 2
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

//切片共享内存 共享内存数据在修改时会对其他数据产生影响
func TestSliceShareMemory(t *testing.T) {
	year := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	q1 := year[3:6]
	q2 := year[5:8]
	t.Log(q1, len(q1), cap(q1)) // [4,5,6] 3 9
	t.Log(q2, len(q2), cap(q2)) // [6,7,8] 3 7
	q2[0] = 13
	t.Log(q1, len(q1), cap(q1)) // [4,5,13] 3 9
	t.Log(q2, len(q2), cap(q2)) // [13,7,8] 3 7

	q1 = append(q1, 14)
	t.Log(q1, len(q1), cap(q1)) // [4,5,13,14] 4 9
	t.Log(q2, len(q2), cap(q2)) // [13,14,8] 3 7
}

//数组和切片的比较
func TestSliceCompare(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 4}
	c := [3]int{1, 2, 3}

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 4}
	s3 := []int{1, 2, 3}

	t.Log(a == b)
	t.Log(a == c)

	t.Log(s1, s2, s3)
	//不能比较 直接报错
	//t.Log(s1 == s2)
	//t.Log(s1 == s3)
}
