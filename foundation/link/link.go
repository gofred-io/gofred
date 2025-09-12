package link

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/basic/anchor"
)

type link struct {
	opts []anchor.Option
}

func New(child application.BaseWidget, opts ...Option) application.BaseWidget {
	link := &link{}

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

	return anchor.New(child, link.opts...)
}
