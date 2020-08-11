package transitions

func Police() []string {
	duration := 300
	brightness := 100
	return []string {
		NewRGBTransition(Red, duration, brightness),
		NewRGBTransition(Blue, duration, brightness),
	}
}

func Police2() []string {
	duration := 200
	brightness := 100
	return []string {
		NewRGBTransition(Red, duration, brightness),
		NewRGBTransition(Red, duration, 1),
		NewRGBTransition(Red, duration, brightness),
		NewSleepTransition(duration),
		NewRGBTransition(Blue, duration, brightness),
		NewRGBTransition(Blue, duration, 1),
		NewRGBTransition(Blue, duration, brightness),
		NewSleepTransition(duration),
	}
}

func StrobeColor() []string {
	duration := 50
	brightness := 100
	return []string {
		NewHSVTransition(140, 100, duration, brightness),
		NewHSVTransition(60, 100, duration, brightness),
		NewHSVTransition(330, 100, duration, brightness),
		NewHSVTransition(0, 100, duration, brightness),
		NewHSVTransition(173, 100, duration, brightness),
		NewHSVTransition(30, 100, duration, brightness),
	}
}

func Alarm() []string {
	duration := 250
	brightness := 100

	return []string{
		NewHSVTransition(0, 100, duration, brightness),
		NewHSVTransition(0, 100, duration, 60),
	}
}