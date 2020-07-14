package reflect_test

import (
	"os"
	"strconv"
	"testing"
)

type Stringer interface {
	String() string
}

type Celsius float64

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + "°C"
}

type Day int

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (d Day) String() string {
	return dayName[d]
}

func print(args ...interface{}) {
	for i, arg := range args {
		if i > 0 {
			_, _ = os.Stdout.WriteString(" ")
		}
		switch a := arg.(type) {
		case Stringer:
			_, _ = os.Stdout.WriteString(a.String())
		case int:
			_, _ = os.Stdout.WriteString(strconv.Itoa(a))
		case string:
			_, _ = os.Stdout.WriteString(a)
		default:
			_, _ = os.Stdout.WriteString("???")
		}
	}
}

func TestStr(t *testing.T) {
	print(Day(1), "was", Celsius(18.32))
}
