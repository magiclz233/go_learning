package tempconv

func CToF(c Celsius) Fahrenheit {
	v := c*9/5 + 32
	return Fahrenheit(v)
}

func FToC(f Fahrenheit) Celsius {
	v := (f - 32) * 5 / 9
	return Celsius(v)
}
