package grid

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/theme"
)

type grid struct {
	opts []div.Option
}

func New(children []application.BaseWidget, opts ...Option) application.BaseWidget {
	g := &grid{}

	opts = append(
		opts,
		display(theme.DisplayTypeGrid),
	)

	for _, option := range opts {
		option(g)
	}

	return div.New(
		children,
		g.opts...,
	)
}
