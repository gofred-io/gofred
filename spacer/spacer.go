package spacer

import (
	"github.com/gofred-io/gofred/container"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

func New(opts ...options.Options) widget.BaseWidget {
	if len(opts) == 0 {
		opts = append(opts, options.Flex(1))
	}

	return container.New(
		widget.Nil,
		opts...,
	)
}
