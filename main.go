package main

import (
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/fsmiamoto/benri/directory"
	"github.com/fsmiamoto/benri/git"
)

type Module struct {
	source       func() string
	colorOptions []color.Attribute
}

var modules = []*Module{
	{directory.CurrentWorking, []color.Attribute{color.FgBlue, color.Bold}},
	{git.CurrentBranch, []color.Attribute{color.FgMagenta, color.Bold}},
	{git.Status, []color.Attribute{color.FgRed, color.Bold}},
}

func main() {
	p, err := buildPrompt(modules)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write([]byte(p + "\n"))
}

func buildPrompt(mods []*Module) (string, error) {
	prompt := make([]string, len(mods))

	for i := range mods {
		src := mods[i].source()
		prompt[i] = color.New(mods[i].colorOptions...).Sprint(src)
	}

	return strings.Join(prompt, " "), nil
}
