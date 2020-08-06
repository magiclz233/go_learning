package tempconv

import "fmt"

//摄氏度
type Celsius float64

//华氏摄氏度
type Fahrenheit float64

const (
	AbsoluteZeroC = -273.15
	FreezingC     = 0
	BoilingC      = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°C", f)
}
