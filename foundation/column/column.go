package column

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type column struct {
	*widget.BaseWidget
	opts []div.Option
}

func New(children []widget.Widget, opts ...Option) *column {
	c := &column{
		BaseWidget: &widget.BaseWidget{},
	}

	opts = append(
		opts,
		display(options.DisplayTypeFlex),
		flexDirection(options.FlexDirectionTypeColumn),
	)

	for _, option := range opts {
		option(c)
	}

	c.Extend(
		div.New(
			children,
			c.opts...,
		),
	)
	return c
}
