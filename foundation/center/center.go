package center

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type center struct {
	*widget.BaseWidget
	opts []div.Option
}

func New(child widget.Widget, opts ...Option) *center {
	c := &center{
		BaseWidget: &widget.BaseWidget{},
	}

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

	c.Extend(
		div.New(
			[]widget.Widget{child},
			c.opts...,
		),
	)

	return c
}
