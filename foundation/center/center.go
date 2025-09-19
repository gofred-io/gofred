package center

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/theme"
)

type center struct {
	opts []div.Option
}

func New(child application.BaseWidget, opts ...Option) application.BaseWidget {
	c := &center{}

	opts = append(
		opts,
		display(theme.DisplayTypeFlex),
		flex(1),
		alignItems(theme.AxisAlignmentTypeCenter),
		justifyContent(theme.AxisAlignmentTypeCenter),
	)

	for _, option := range opts {
		option(c)
	}

	return div.New(
		[]application.BaseWidget{child},
		c.opts...,
	)
}
