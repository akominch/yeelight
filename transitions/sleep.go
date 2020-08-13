package transitions

import (
	"fmt"
	"github.com/akominch/yeelight/utils"
	"strconv"
)

type Sleep struct {
	duration int
	mode int
}

func NewSleepTransition(duration int) *Sleep {
	return &Sleep{
		duration: utils.GetDurationValue(duration),
		mode:     7,
	}
}

func (t *Sleep) AsYeelightParams() string {
	return fmt.Sprintf("%s,%s,1,2", strconv.Itoa(t.duration), strconv.Itoa(t.mode))
}

func (t *Sleep) ChangeDuration(duration int) {
	t.duration = utils.GetDurationValue(duration)
}