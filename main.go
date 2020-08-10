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
	{Content: directory.CurrentWorking, After: " ", ColorOptions: []color.Attribute{color.FgBlue, color.Bold}},
	{Content: git.CurrentBranch, After: " ", ColorOptions: []color.Attribute{color.FgMagenta, color.Bold}},
	{Content: git.HasModified("!"), ColorOptions: []color.Attribute{color.FgHiGreen, color.Bold}},
	{Content: git.HasUntracked("?"), After: " ", ColorOptions: []color.Attribute{color.FgHiBlue, color.Bold}},
	{Content: git.HasStashed("$"), ColorOptions: []color.Attribute{color.FgHiMagenta, color.Bold}},
	{Content: git.HasStaged("+"), ColorOptions: []color.Attribute{color.FgHiRed, color.Bold}},
	{Content: duration.String, ColorOptions: []color.Attribute{color.FgYellow, color.Bold}},
}

func main() {
	fmt.Println(prompt.New(modules))
}
