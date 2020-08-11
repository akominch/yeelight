package transitions

import (
	"fmt"
	"log"
	"strconv"
)

func NewSleepTransition(duration int) string {
	if duration < 50 {
		log.Fatalln("duration minimum 50")
	}

	mode := strconv.Itoa(7)

	// Ignored by Yeelight
	value := strconv.Itoa(1)
	brightness := strconv.Itoa(2)

	return fmt.Sprintf("%s,%s,%s,%s", strconv.Itoa(duration), mode, value, brightness)
}
