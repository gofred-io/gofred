package button

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme/theme_style"
)

type Option options.OptionWrapper

func BackgroundColor(backgroundColor string) Option {
	return func() options.Option {
		return options.BackgroundColor(backgroundColor)
	}
}

func BorderRadius(borderRadius int) Option {
	return func() options.Option {
		return options.BorderRadius(borderRadius)
	}
}

func ButtonStyle(buttonStyle theme_style.ButtonStyle) Option {
	return func() options.Option {
		return options.ButtonStyle(buttonStyle)
	}
}

func Class(class string) Option {
	return func() options.Option {
		return options.Class(class)
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

func Label(label string) Option {
	return func() options.Option {
		return options.Label(label)
	}
}

func MaxWidth(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.MaxWidth(opts...)
	}
}

func OnClick(handler options.OnClickHandler) Option {
	return func() options.Option {
		return options.OnClick(handler)
	}
}

func Padding(opts ...breakpoint.BreakpointOptions[spacing.Spacing]) Option {
	return func() options.Option {
		return options.Padding(opts...)
	}
}

func Tooltip(tooltip string) Option {
	return func() options.Option {
		return options.Tooltip(tooltip)
	}
}

func Visible(opts ...breakpoint.BreakpointOptions[bool]) Option {
	return func() options.Option {
		return options.Visible(opts...)
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
