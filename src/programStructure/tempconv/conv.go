package tempconv

// 转换函数

// CtoF把摄氏温度转换为华氏温度
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FtoC把华氏温度转换为摄氏温度
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
