package grid

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type grid struct {
	opts []div.Option
}

func New(children []widget.BaseWidget, opts ...Option) widget.BaseWidget {
	g := &grid{}

	opts = append(
		opts,
		display(options.DisplayTypeGrid),
	)

	for _, option := range opts {
		option(g)
	}

	return div.New(
		children,
		g.opts...,
	)
}
