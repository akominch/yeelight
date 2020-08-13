package utils

func GetBrightnessValue(brightness int) int {
	return GetValue(brightness, 1, 100)
}

func GetSaturationValue(saturation int) int {
	return GetValue(saturation, 1, 100)
}

func GetHueValue(hue int) int {
	return GetValue(hue, 0, 359)
}

func GetDegreesValue(degrees int) int {
	return GetValue(degrees, 1700, 6500)
}

func GetDurationValue(duration int) int {
	return minInt(duration, 50)
}

func GetValue(x, min, max int) int {
	return maxInt(minInt(x, min), max)
}

func minInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}