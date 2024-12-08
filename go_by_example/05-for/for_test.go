package for_test

import "testing"

func TestFor(t *testing.T) {
	i := 1
	for i < 3 {
		t.Logf("i值：%d", i)
		i++
	}
	for i := 0; i < 10; i++ {
		t.Logf("i：%d", i)
	}

	list := make([]int, 5)
	for index, value := range list {
		t.Log(index, value)
	}
}
