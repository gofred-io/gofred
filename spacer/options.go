package spacer

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/options"
)

type Options func() options.Options

func flex(flex int) Options {
	return func() options.Options {
		return options.Flex(flex)
	}
}

func Height(height int) Options {
	return func() options.Options {
		return options.Height(breakpoint.All(height))
	}
}

func Width(width int) Options {
	return func() options.Options {
		return options.Width(breakpoint.All(width))
	}
}
