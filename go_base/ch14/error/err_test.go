package err_test

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n > 100 {
		return nil, LargerThanHundredError
		//return nil, errors.New("n should be not larger than 100")

	}
	if n < 2 {
		return nil, LessThanTwoError
		//return nil, errors.New("n should be not less than 2")

	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	t.Log(GetFibonacci(-10))
	t.Log(GetFibonacci(20))
	t.Log(GetFibonacci(110))

	if v, err := GetFibonacci(5); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less")
		} else if err == LargerThanHundredError {
			fmt.Println("It is larger")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}
