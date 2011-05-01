package clock

import (
	"time"
)

func IsAm() bool {
	localTime := time.LocalTime()
	return localTime.Hour <= 12
}