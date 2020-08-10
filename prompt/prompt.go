package prompt

import (
	"strings"

	"github.com/fatih/color"
)

type Module struct {
	Content      func() string
	Before       string
	After        string
	ColorOptions []color.Attribute
}

func New(modules []*Module) string {
	contents := make([]string, 0, len(modules))
	for i := range modules {
		if content := modules[i].Content(); content != "" {
			content = modules[i].Before + content + modules[i].After
			contents = append(contents, color.New(modules[i].ColorOptions...).Sprint(content))
		}
	}
	return strings.Join(contents, "")
}
