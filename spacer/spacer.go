package spacer

import (
	"github.com/gofred-io/gofred/container"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

func New(opts ...Options) widget.BaseWidget {
	var (
		containerOptions = []options.Options{}
	)

	if len(opts) == 0 {
		opts = append(opts, flex(1))
	}

	for _, option := range opts {
		containerOptions = append(containerOptions, option())
	}

	return container.New(
		widget.Nil,
		containerOptions...,
	)
}
