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
	sources := make([]string, 0, len(modules))
	for i := range modules {
		if source := modules[i].Source(); source != "" {
			sources = append(sources, color.New(modules[i].ColorOptions...).Sprint(source))
		}
	}
	return strings.Join(sources, " ")
}
