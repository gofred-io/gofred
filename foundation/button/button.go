package button

import (
	basicbutton "github.com/gofred-io/gofred/basic/button"
	"github.com/gofred-io/gofred/widget"
)

type button struct {
	*widget.BaseWidget
	opts []basicbutton.Option
}

func New(child widget.Widget, opts ...Option) *button {
	var (
		b = &button{
			BaseWidget: &widget.BaseWidget{},
		}
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

	b.Extend(
		basicbutton.New(
			child,
			b.opts...,
		),
	)

	return b
}
