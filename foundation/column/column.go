package column

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type column struct {
	opts []div.Option
}

func New(children []widget.BaseWidget, opts ...Option) widget.BaseWidget {
	c := &column{}

	opts = append(
		opts,
		display(options.DisplayTypeFlex),
		flexDirection(options.FlexDirectionTypeColumn),
	)

	for _, option := range opts {
		option(c)
	}

	return div.New(
		children,
		c.opts...,
	)
}
