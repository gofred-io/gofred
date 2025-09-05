package grid

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type grid struct {
	*widget.BaseWidget
	opts []div.Option
}

func New(children []widget.Widget, opts ...Option) *grid {
	g := &grid{
		BaseWidget: &widget.BaseWidget{},
	}

	opts = append(
		opts,
		display(options.DisplayTypeGrid),
	)

	for _, option := range opts {
		option(g)
	}

	g.Extend(
		div.New(
			children,
			g.opts...,
		),
	)
	return g
}
