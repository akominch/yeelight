package transitions

import (
	"fmt"
	"github.com/akominch/yeelight/utils"
	"strconv"
)

type Temperature struct {
	degrees int
	duration int
	brightness int
	mode int
}

func NewTemperatureTransition(degrees int, duration int, brightness int) *Temperature {
	return &Temperature{
		degrees:    utils.GetDegreesValue(degrees),
		duration:   utils.GetDurationValue(duration),
		brightness: utils.GetBrightnessValue(brightness),
		mode:       2,
	}
}

func (t *Temperature) AsYeelightParams() string {
	return fmt.Sprintf("%s,%s,%s,%s", strconv.Itoa(t.duration), strconv.Itoa(t.mode), strconv.Itoa(t.degrees), strconv.Itoa(t.brightness))
}

func (t *Temperature) ChangeDuration(duration int) {
	t.duration = utils.GetDurationValue(duration)
}