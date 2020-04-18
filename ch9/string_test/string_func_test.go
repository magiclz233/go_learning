package string

import (
	"strings"
	"testing"
)

//string的函数
func TestStringFunc(t *testing.T) {
	s := "A,B,C,D,E,F"
	split := strings.Split(s, ",")
	for _,v := range split{
		t.Log(v)
	}

	t.Log(strings.Join(split,"-"))
}
