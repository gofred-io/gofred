package path

import (
	"github.com/gofred-io/gofred/options"
)

type Option options.OptionWrapper

func d(data string) Option {
	return func() options.Option {
		return options.D(data)
	}
}

func Class(class string) Option {
	return func() options.Option {
		return options.Class(class)
	}
}

func Fill(fill string) Option {
	return func() options.Option {
		return options.Fill(fill)
	}
}

func ID(id string) Option {
	return func() options.Option {
		return options.ID(id)
	}
}
