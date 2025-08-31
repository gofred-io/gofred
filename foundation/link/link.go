package link

import (
	"github.com/gofred-io/gofred/basic/anchor"
	"github.com/gofred-io/gofred/widget"
)

type link struct {
	opts []anchor.Option
}

func New(child widget.BaseWidget, opts ...Option) widget.BaseWidget {
	link := &link{}

	opts = append(
		opts,
		Class("gf-link"),
	)

	for _, option := range opts {
		option(link)
	}

	return anchor.New(child, link.opts...)
}
