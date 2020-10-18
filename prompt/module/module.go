package module

import "github.com/fatih/color"

type Content func() string

type Module struct {
	Content      Content
	Before       string
	After        string
	ColorOptions []color.Attribute
}
