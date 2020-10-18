package clock

import "time"

func With24H() string {
	return time.Now().Format("15:04")
}

func With12H() string {
	return time.Now().Format("3:04PM")
}

func WithFormat(fmt string) string {
	return time.Now().Format(fmt)
}
