package link

import (
	"github.com/gofred-io/gofred/basic/anchor"
	"github.com/gofred-io/gofred/widget"
)

type link struct {
	*widget.BaseWidget
	opts []anchor.Option
}

func New(child widget.Widget, opts ...Option) *link {
	link := &link{
		BaseWidget: &widget.BaseWidget{},
	}

	defaultOpts := []Option{
		Class("gf-link"),
	}

	opts = append(
		defaultOpts,
		opts...,
	)

	for _, option := range opts {
		option(link)
	}

	link.Extend(
		anchor.New(child, link.opts...),
	)
	return link
}
