package container

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type container struct {
	*widget.BaseWidget
	opts []div.Option
}

func New(child widget.Widget, opts ...Option) *container {
	var (
		c = &container{
			BaseWidget: &widget.BaseWidget{},
		}
	)

	opts = append(
		opts,
		display(options.DisplayTypeFlex),
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
