package transitions

import (
	"fmt"
	color2 "github.com/akominch/yeelight/color"
	"github.com/akominch/yeelight/utils"
	"image/color"
	"strconv"
)

var (
	Red = color.RGBA{R: 255, A: 255}
	Blue = color.RGBA{B: 255, A: 255}
	Green = color.RGBA{G: 255, A: 255}
)

type RGB struct {
	color color.RGBA
	duration int
	brightness int
	mode int
	yeelightValue int
}

func NewRGBTransition(color color.RGBA, duration int, brightness int) *RGB {
	value := color2.RGBToYeelight(color)

	return &RGB{
		color:         color,
		brightness:    utils.GetBrightnessValue(brightness),
		duration:      utils.GetDurationValue(duration),
		yeelightValue: value,
		mode:          1,
	}
}

func (t *RGB) AsYeelightParams() string {
	return fmt.Sprintf("%s,%s,%s,%s", strconv.Itoa(t.duration), strconv.Itoa(t.mode), strconv.Itoa(t.yeelightValue), strconv.Itoa(t.brightness))
}

func (t *RGB) ChangeDuration(duration int) {
	t.duration = utils.GetDurationValue(duration)
}

func (t *RGB) ChangeColor(color color.RGBA) {
	value := color2.RGBToYeelight(color)

	t.color = color
	t.yeelightValue = value
}