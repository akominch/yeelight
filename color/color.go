package color

import (
	"image/color"
)

func RGBToYeelight(color color.RGBA) int {
	r := int(color.R)
	g := int(color.G)
	b := int(color.B)

	return r * 65536 + g * 256 + b
}