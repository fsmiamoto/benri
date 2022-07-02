package text

func FromString(s string) func() string {
	return func() string {
		return s
	}
}
