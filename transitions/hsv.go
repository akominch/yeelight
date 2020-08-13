package transitions

import (
	"fmt"
	"github.com/akominch/yeelight/color"
	"github.com/akominch/yeelight/utils"
	"github.com/lucasb-eyer/go-colorful"
	color2 "image/color"
	"strconv"
)

type HSV struct {
	hue int
	saturation int
	duration int
	brightness int
	mode int
	yeelightValue int
}

func NewHSVTransition(hue int, saturation int, duration int, brightness int) *HSV {
	value := color.RGBToYeelight(convertHSVToRGB(hue, saturation))

	return &HSV{
		hue:           utils.GetHueValue(hue),
		saturation:    utils.GetSaturationValue(saturation),
		duration:      utils.GetDurationValue(duration),
		brightness:    utils.GetBrightnessValue(brightness),
		yeelightValue: value,
		mode:          1,
	}
}

func (t *HSV) AsYeelightParams() string {
	return fmt.Sprintf("%s,%s,%s,%s", strconv.Itoa(t.duration), strconv.Itoa(t.mode), strconv.Itoa(t.yeelightValue), strconv.Itoa(t.brightness))
}

func (t *HSV) ChangeColor(hue int, saturation int) {
	value := color.RGBToYeelight(convertHSVToRGB(hue, saturation))

	t.hue = hue
	t.saturation = saturation
	t.yeelightValue = value
}

func (t *HSV) ChangeDuration(duration int) {
	t.duration = utils.GetDurationValue(duration)
}

func convertHSVToRGB(hue int, saturation int) color2.RGBA {
	color := colorful.Hsv(float64(hue), float64(saturation), 1)
	return color2.RGBA{
		R: uint8(color.R),
		G: uint8(color.G),
		B: uint8(color.B),
	}
}