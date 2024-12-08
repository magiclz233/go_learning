package loop

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

var ear = "AD"

func TestWhileLoop1(t *testing.T) {
	year := 2018
	month := rand.Intn(12) + 1
	daysInMonth := 31
	switch month {
	case 2:
		if isLeapYear(year) {
			daysInMonth = 29
		} else {
			daysInMonth = 28
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	day := rand.Intn(daysInMonth) + 1
	fmt.Println(ear, year, month, day)

	v := 42
	if v >= 0 && v <= math.MaxUint8 {
		v8 := uint8(v)
		fmt.Println("v8:", v8)
	}
	testString := []string{"1", "0", "true", "false", "3"}

	for _, str := range testString {
		result, err := strTransBol(str)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s - %t \n", str, result)
		}
	}
}

func isLeapYear(year int) bool {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

func strTransBol(str string) (bool, error) {
	switch str {
	case "true", "yes", "1":
		return true, nil
	case "no", "false", "0":
		return false, nil
	}
	return false, fmt.Errorf("请传入正确的字符串： %s", str)
}
