package container

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type container struct {
	opts []div.Option
}

func New(child widget.BaseWidget, opts ...Option) widget.BaseWidget {
	var (
		c = &container{}
	)

	opts = append(
		opts,
		display(options.DisplayTypeFlex),
	)

	for _, option := range opts {
		option(c)
	}

	return div.New(
		[]widget.BaseWidget{child},
		c.opts...,
	)
}
