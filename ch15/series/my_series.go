package series

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("series init1")
}

func init() {
	fmt.Println("series init2")
}

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

func square(n int) int {
	return n * n
}

func Square(n int) int {
	return n * n
}
