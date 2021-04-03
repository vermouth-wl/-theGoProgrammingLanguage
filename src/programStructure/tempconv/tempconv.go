package tempconv

import "fmt"

// 该包负责摄氏温度与华氏温度的转换

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g℉", f)
}
