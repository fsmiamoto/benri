package module

import "github.com/fatih/color"

type Content func() string

type Option func(*Module)

type Module struct {
	Content      Content
	Before       string
	After        string
	ColorOptions []color.Attribute
}

func New(c Content, opts ...Option) *Module {
	m := &Module{Content: c}
	for _, option := range opts {
		option(m)
	}
	return m
}

func Before(b string) Option {
	return func(m *Module) {
		m.Before = b
	}
}

func After(a string) Option {
	return func(m *Module) {
		m.After = a
	}
}

func ColorOptions(attrs ...color.Attribute) Option {
	return func(m *Module) {
		m.ColorOptions = attrs
	}
}
