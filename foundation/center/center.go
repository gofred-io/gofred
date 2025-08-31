package center

import (
	"github.com/gofred-io/gofred/div"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type center struct {
	opts []div.Option
}

func New(child widget.BaseWidget, opts ...Option) widget.BaseWidget {
	c := &center{}

	opts = append(
		opts,
		display(options.DisplayTypeFlex),
		flex(1),
		alignItems(options.AxisAlignmentTypeCenter),
		justifyContent(options.AxisAlignmentTypeCenter),
	)

	for _, option := range opts {
		option(c)
	}

	return div.New(
		[]widget.BaseWidget{child},
		c.opts...,
	)
}
