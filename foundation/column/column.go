package column

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/theme"
)

type column struct {
	opts []div.Option
}

func New(children []application.BaseWidget, opts ...Option) application.BaseWidget {
	c := &column{}

	opts = append(
		opts,
		display(theme.DisplayTypeFlex),
		flexDirection(theme.FlexDirectionTypeColumn),
	)

	for _, option := range opts {
		option(c)
	}

	return div.New(
		children,
		c.opts...,
	)
}
