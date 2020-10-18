package prompt

import (
	"strings"

	"github.com/fatih/color"
	m "github.com/fsmiamoto/benri/prompt/module"
)

func New(modules []*m.Module) string {
	contents := make([]string, 0, len(modules))
	for i := range modules {
		if content := modules[i].Content(); content != "" {
			content = modules[i].Before + content + modules[i].After
			contents = append(contents, color.New(modules[i].ColorOptions...).Sprint(content))
		}
	}
	return strings.Join(contents, "")
}
