package prompt

import (
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
	m "github.com/fsmiamoto/benri/prompt/module"
	"golang.org/x/crypto/ssh/terminal"
)

func New(modules []*m.Module) string {
	p, _ := buildFromModules(modules)
	return p
}

func NewWithLeftAndRight(left []*m.Module, right []*m.Module) string {
	l, ll := buildFromModules(left)
	r, rl := buildFromModules(right)

	whitespace := getTerminalWidth() - ll - rl
	w := strings.Repeat(" ", whitespace)

	return l + w + r
}

func buildFromModules(modules []*m.Module) (string, int) {
	contents := make([]string, 0, len(modules))
	length := 0
	for i := range modules {
		if content := modules[i].Content(); content != "" {
			content = modules[i].Before + content + modules[i].After
			length += utf8.RuneCountInString(content)
			contents = append(contents, color.New(modules[i].ColorOptions...).Sprint(content))
		}
	}
	return strings.Join(contents, ""), length
}

func getTerminalWidth() int {
	width, _, _ := terminal.GetSize(0)
	return width
}
