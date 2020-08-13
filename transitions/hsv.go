package transitions

import (
	"fmt"
	"github.com/akominch/yeelight"
	"github.com/lucasb-eyer/go-colorful"
	color2 "image/color"
	"log"
	"strconv"
)

func NewHSVTransition(hue int, saturation int, duration int, brightness int) string {
	if hue < 0 || hue > 359 {
		log.Fatalln("The color hue to transition to (0-359)")
	}
	if saturation < 0 || saturation > 100 {
		log.Fatalln("The color saturation to transition to (0-100)")
	}
	if brightness < 0 || brightness > 100 {
		log.Fatalln("The color saturation to transition to (0-100)")
	}
	if duration < 50 {
		log.Fatalln("duration minimum 50")
	}

	color := colorful.Hsv(float64(hue), float64(saturation), 1)
	rgb := color2.RGBA{
		R: uint8(color.R),
		G: uint8(color.G),
		B: uint8(color.B),
	}
	value := yeelight.RGBToYeelight(rgb)

	return fmt.Sprintf("%s,1,%s,%s", strconv.Itoa(duration), strconv.Itoa(value), strconv.Itoa(brightness))
}
