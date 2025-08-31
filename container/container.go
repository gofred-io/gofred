package container

import (
	"github.com/gofred-io/gofred/div"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type container struct {
	opts []div.Option
}

func New(child widget.BaseWidget, opts ...Option) widget.BaseWidget {
	var (
		c        = &container{}
		children = []widget.BaseWidget{}
	)

	opts = append(
		opts,
		display(options.DisplayTypeFlex),
	)

	for _, option := range opts {
		option(c)
	}

	if !child.Equal(widget.Nil) {
		children = append(children, child)
	}

	return div.New(
		children,
		c.opts...,
	)
}
