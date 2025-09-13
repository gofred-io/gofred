package pre

import (
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/theme/theme_style"
)

type Option options.OptionWrapper

func Class(class string) Option {
	return func() options.Option {
		return options.Class(class)
	}
}

func ID(id string) Option {
	return func() options.Option {
		return options.ID(id)
	}
}

func TextStyle(textStyle theme_style.TextStyle) Option {
	return func() options.Option {
		return options.TextStyle(textStyle)
	}
}
