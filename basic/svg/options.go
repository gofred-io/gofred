package svg

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/options"
)

type Option options.OptionWrapper

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

func Height(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.Height(opts...)
	}
}

func ID(id string) Option {
	return func() options.Option {
		return options.ID(id)
	}
}

func MaxWidth(maxWidth int) Option {
	return func() options.Option {
		return options.MaxWidth(maxWidth)
	}
}

func UserSelect(userSelect options.UserSelectType) Option {
	return func() options.Option {
		return options.UserSelect(userSelect)
	}
}

func Width(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.Width(opts...)
	}
}

func WidthP(opts ...breakpoint.BreakpointOptions[float64]) Option {
	return func() options.Option {
		return options.WidthP(opts...)
	}
}
