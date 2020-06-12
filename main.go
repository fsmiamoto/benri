package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/fsmiamoto/benri/modules/directory"
	"github.com/fsmiamoto/benri/modules/git"
	"github.com/fsmiamoto/benri/prompt"
)

var modules = []*prompt.Module{
	{directory.CurrentWorking, []color.Attribute{color.FgBlue, color.Bold}},
	{git.CurrentBranch, []color.Attribute{color.FgMagenta, color.Bold}},
	{git.Status, []color.Attribute{color.FgCyan, color.Bold}},
}

func main() {
	fmt.Print(prompt.New(modules), "\n", "$ ")
}
