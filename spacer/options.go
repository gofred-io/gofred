package spacer

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/div"
)

type Option func(spacer *spacer)

func flex(flex int) Option {
	return func(spacer *spacer) {
		spacer.opts = append(spacer.opts, div.Flex(flex))
	}
}

func Height(height int) Option {
	return func(spacer *spacer) {
		spacer.opts = append(spacer.opts, div.Height(breakpoint.All(height)))
	}
}

func Width(width int) Option {
	return func(spacer *spacer) {
		spacer.opts = append(spacer.opts, div.Width(breakpoint.All(width)))
	}
}
