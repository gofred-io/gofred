package container

import (
	"github.com/gofred-io/gofred/div"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

func New(child widget.BaseWidget, opts ...options.Options) widget.BaseWidget {
	var (
		children = []widget.BaseWidget{}
	)

	opts = append(
		opts,
		options.Display(options.DisplayTypeFlex),
	)

	if !child.Equal(widget.Nil) {
		children = append(children, child)
	}

	return div.New(
		children,
		opts...,
	)
}
