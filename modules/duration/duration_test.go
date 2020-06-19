package duration

import "testing"

func TestFormat(t *testing.T) {
	tests := []struct {
		name     string
		duration int64
		expect   string
	}{
		{"Just seconds", 3, "3s"},
		{"Minutes and seconds", 120 + 42, "2m42s"},
		{"Hours, minutes and seconds", 5*3600 + 17*60 + 15, "5h17m15s"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := format(tt.duration)

			if got != tt.expect {
				t.Errorf("Got %v but expected %v", got, tt.expect)
			}
		})
	}
}
