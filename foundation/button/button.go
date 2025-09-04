package button

import (
	basicbutton "github.com/gofred-io/gofred/basic/button"
	"github.com/gofred-io/gofred/widget"
)

type button struct {
	opts []basicbutton.Option
}

func New(child widget.BaseWidget, opts ...Option) widget.BaseWidget {
	var (
		b = &button{}
	)

	defaultOpts := []Option{
		Class("gf-button"),
	}

	opts = append(
		defaultOpts,
		opts...,
	)

	for _, option := range opts {
		option(b)
	}

	return basicbutton.New(
		child,
		b.opts...,
	)
}
