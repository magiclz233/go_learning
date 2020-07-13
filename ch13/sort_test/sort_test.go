package sort_test

import (
	"fmt"
	"go_learning/ch13/sort"
	"testing"
)

func intSort() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := sort.IntArray(data)
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("sort int array failed!")
	}
	fmt.Printf("sorted int array: %v\n", a)
}

func stringSort() {
	data := []string{"monday", "friday", "tuesday", "wednesday", "sunday", "thursday", "", "saturday"}
	a := sort.StringArray(data)
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("sort string array failed")
	}
	fmt.Printf("sorted string array: %v\n", a)
}

type day struct {
	num       int
	shortName string
	longName  string
}

type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int {
	return len(p.data)
}

func (p *dayArray) Less(i, j int) bool {
	return p.data[i].num < p.data[j].num
}

func (p *dayArray) Swap(i, j int) {
	p.data[i], p.data[j] = p.data[j], p.data[i]
}

func days() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	a := dayArray{data}
	sort.Sort(&a)
	if !sort.IsSorted(&a) {
		panic("*dayArray sorted failed!")
	}

	for k, v := range data {
		fmt.Printf("k: %d, v: %s\n", k, v.longName)
	}
}

func TestSort(t *testing.T) {
	intSort()
	stringSort()
	days()
}
