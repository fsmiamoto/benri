package duration

import (
	"os"
	"strconv"
	"time"
)

// Durations less than MinDuration are not shown
const MinDuration = 3

func Seconds() string {
	epoch, err := strconv.Atoi(os.Getenv("BENRI_PREEXEC"))
	if err != nil {
		return ""
	}

	duration := time.Now().Unix() - int64(epoch)
	if duration < MinDuration {
		return ""
	}

	return strconv.FormatInt(duration, 10) + "s"
}
