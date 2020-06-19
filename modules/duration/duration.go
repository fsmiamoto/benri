package duration

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// Durations less than MinDuration are not shown
const MinDuration = 3 // seconds

const Day = 24 * Hour
const Hour = int64(time.Hour / 1E9)
const Min = int64(time.Minute / 1E9)
const Sec = int64(time.Second / 1E9)

func String() string {
	timestamp, err := strconv.Atoi(os.Getenv("BENRI_PREEXEC"))
	if err != nil {
		return ""
	}

	duration := time.Now().Unix() - int64(timestamp)
	if duration < MinDuration {
		return ""
	}

	return format(duration)
}

// format returns a string with human readable format for
// a duration in seconds
func format(dur int64) string {
	sections := []string{"", "", ""}

	if dur/Hour != 0 {
		hours := dur / Hour
		sections[0] = strconv.FormatInt(hours, 10) + "h"
		dur -= hours * Hour
	}

	if dur/Min != 0 {
		mins := dur / Min
		sections[1] = strconv.FormatInt(mins, 10) + "m"
		dur -= mins * Min
	}

	if dur/Sec != 0 {
		secs := dur / Sec
		sections[2] = strconv.FormatInt(secs, 10) + "s"
		dur -= secs * Sec
	}

	return strings.Join(sections, "")
}
