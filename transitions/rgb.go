package transitions

import (
	"fmt"
	"github.com/akominch/yeelight"
	"image/color"
	"log"
	"strconv"
)

var (
	Red = color.RGBA{R: 255, A: 255}
	Blue = color.RGBA{B: 255, A: 255}
	Green = color.RGBA{G: 255, A: 255}
)

func NewRGBTransition(color color.RGBA, duration int, brightness int) string {
	if brightness < 0 || brightness > 100 {
		log.Fatalln("The color saturation to transition to (0-100)")
	}
	if duration < 50 {
		log.Fatalln("duration minimum 50")
	}

	value := yeelight.RGBToYeelight(color)

	return fmt.Sprintf("%s,%s,%s,%s", strconv.Itoa(duration), strconv.Itoa(1), strconv.Itoa(value), strconv.Itoa(brightness))
}
