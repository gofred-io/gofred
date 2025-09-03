package pre

import (
	"github.com/gofred-io/gofred/options"
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
