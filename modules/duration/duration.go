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
	return format(int64(timestamp))
}

// format returns a string with human readable format for
// a Unix Time
func format(timestamp int64) string {
	elapsed := time.Now().Unix() - int64(timestamp)
	if elapsed < MinDuration {
		return ""
	}

	sections := []string{"", "", ""}

	if elapsed/Hour != 0 {
		hours := elapsed / Hour
		sections[0] = strconv.FormatInt(hours, 10) + "h"
		elapsed -= hours * Hour
	}

	if elapsed/Min != 0 {
		mins := elapsed / Min
		sections[1] = strconv.FormatInt(mins, 10) + "m"
		elapsed -= mins * Min
	}

	if elapsed/Sec != 0 {
		secs := elapsed / Sec
		sections[2] = strconv.FormatInt(secs, 10) + "s"
		elapsed -= secs * Sec
	}

	return strings.Join(sections, "")
}
