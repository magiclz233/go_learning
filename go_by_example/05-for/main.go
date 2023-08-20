package main

import (
	"testing"
)

func TestFor(t *testing.T) {
	i := 1
	for i < 3 {
		t.Logf("i值：%d", i)
		i++
	}
}
