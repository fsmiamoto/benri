package prompt

import "testing"

func TestNew(t *testing.T) {
	type args struct {
		modules []*Module
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic prompt",
			args: args{
				modules: []*Module{
					{
						Content: func() string { return "user@computer" },
						After:   " ",
						Before:  "! ",
					},
					{
						Content: func() string { return "~/projects" },
						After:   " ",
						Before:  "at ",
					},
					{
						Content: func() string { return "3s" },
					},
				},
			},
			want: "! user@computer at ~/projects 3s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.modules); got != tt.want {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
