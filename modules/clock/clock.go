package clock

import "time"

func With24H() string {
    return time.Now().Format("15:04:05")
}

func With12H() string {
    return time.Now().Format("3:04:05PM")
}

func WithFormat(fmt string) func() string {
	return func() string {
		return time.Now().Format(fmt)
	}

}
