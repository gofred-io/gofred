package grid

import (
	"github.com/gofred-io/gofred/div"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

func New(children []widget.BaseWidget, opts ...options.Options) widget.BaseWidget {
	opts = append(
		opts,
		options.Display(options.DisplayTypeGrid),
	)

	return div.New(
		children,
		opts...,
	)
}
