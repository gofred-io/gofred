package button

import (
	"github.com/gofred-io/gofred/application"
	basicbutton "github.com/gofred-io/gofred/basic/button"
)

type button struct {
	opts []basicbutton.Option
}

func New(child application.BaseWidget, opts ...Option) application.BaseWidget {
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
