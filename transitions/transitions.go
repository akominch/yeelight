package transitions

type Transition interface {
	AsYeelightParams() string
	ChangeDuration(d int)
}

func Police() []Transition {
	duration := 300
	brightness := 100
	return []Transition {
		NewRGBTransition(Red, duration, brightness),
		NewRGBTransition(Blue, duration, brightness),
	}
}

func Police2() []Transition {
	duration := 100
	brightness := 100
	return []Transition {
		NewRGBTransition(Red, duration, brightness),
		NewRGBTransition(Red, duration, 1),
		NewRGBTransition(Red, duration, brightness),
		NewRGBTransition(Red, duration, 1),
		NewRGBTransition(Red, duration, brightness),
		NewSleepTransition(duration),
		NewRGBTransition(Blue, duration, brightness),
		NewRGBTransition(Blue, duration, 1),
		NewRGBTransition(Blue, duration, brightness),
		NewRGBTransition(Blue, duration, 1),
		NewRGBTransition(Blue, duration, brightness),
		NewSleepTransition(duration),
	}
}

func StrobeColor() []Transition {
	duration := 50
	brightness := 100
	return []Transition {
		NewHSVTransition(140, 100, duration, brightness),
		NewHSVTransition(60, 100, duration, brightness),
		NewHSVTransition(330, 100, duration, brightness),
		NewHSVTransition(0, 100, duration, brightness),
		NewHSVTransition(173, 100, duration, brightness),
		NewHSVTransition(30, 100, duration, brightness),
	}
}

func Alarm() []Transition {
	duration := 250
	brightness := 100

	return []Transition{
		NewHSVTransition(0, 100, duration, brightness),
		NewHSVTransition(0, 100, duration, 60),
	}
}