package prompt

import (
	"strings"

	"github.com/fatih/color"
)

type Module struct {
	Source       func() string
	ColorOptions []color.Attribute
}

func New(modules []*Module) string {
	parts := make([]string, len(modules))
	for i := range modules {
		parts[i] = color.New(modules[i].ColorOptions...).Sprint(modules[i].Source())
	}
	return strings.Join(parts, " ")
}
