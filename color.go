package yeelight

import (
	"image/color"
	"log"
)

func RGBToYeelight(color color.RGBA) int {
	if !IsValid(color) {
		log.Fatalln("Not valid color")
	}

	r := int(color.R)
	g := int(color.G)
	b := int(color.B)

	return r * 65536 + g * 256 + b
}

func IsValid(color color.RGBA) bool {
	switch {
	case color.R < 0 || color.R > 255:
	case color.G < 0 || color.G > 255:
	case color.B < 0 || color.B > 255:
		return false
	}

	return true
}