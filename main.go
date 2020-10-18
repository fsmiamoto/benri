package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/fsmiamoto/benri/modules/directory"
	"github.com/fsmiamoto/benri/modules/duration"
	"github.com/fsmiamoto/benri/modules/git"
	p "github.com/fsmiamoto/benri/prompt"
	m "github.com/fsmiamoto/benri/prompt/module"
)

var modules = []*m.Module{
	m.New(directory.CurrentWorking, m.After(" "), m.ColorOptions(color.FgBlue, color.Bold)),
	m.New(git.CurrentBranch, m.After(" "), m.ColorOptions(color.FgMagenta, color.Bold)),
	m.New(git.HasModified("!"), m.ColorOptions(color.FgHiGreen, color.Bold)),
	m.New(git.HasUntracked("?"), m.ColorOptions(color.FgHiBlue, color.Bold)),
	m.New(git.HasStaged("+"), m.ColorOptions(color.FgHiRed, color.Bold)),
	m.New(git.HasStashed("$"), m.ColorOptions(color.FgHiMagenta, color.Bold)),
	m.New(duration.String, m.Before(" "), m.ColorOptions(color.FgYellow, color.Bold)),
}

func main() {
	fmt.Println(p.New(modules))
}
