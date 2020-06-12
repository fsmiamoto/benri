package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/fsmiamoto/benri/modules/directory"
	"github.com/fsmiamoto/benri/modules/duration"
	"github.com/fsmiamoto/benri/modules/git"
	"github.com/fsmiamoto/benri/prompt"
)

var modules = []*prompt.Module{
	{Source: directory.CurrentWorking, ColorOptions: []color.Attribute{color.FgBlue, color.Bold}},
	{Source: git.CurrentBranch, ColorOptions: []color.Attribute{color.FgMagenta, color.Bold}},
	{Source: duration.Seconds, ColorOptions: []color.Attribute{color.FgYellow, color.Bold}},
}

func main() {
	fmt.Println(prompt.New(modules))
}
