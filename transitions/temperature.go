package transitions

import (
	"fmt"
	"log"
	"strconv"
)

func NewTemperatureTransition(degrees int, duration int, brightness int) string {
	if degrees < 1700 || degrees > 6500 {
		log.Fatalln("The degrees to set the color temperature to (1700-6500)")
	}
	if brightness < 1 || brightness > 100 {
		log.Fatalln("The color saturation to transition to (1-100)")
	}
	if duration < 50 {
		log.Fatalln("duration minimum 50")
	}

	mode := strconv.Itoa(2)
	value := strconv.Itoa(degrees)

	return fmt.Sprintf("%s,%s,%s,%s", strconv.Itoa(duration), mode, value, strconv.Itoa(brightness))
}