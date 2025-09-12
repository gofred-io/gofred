package container

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

type Option func(container *container)

func display(display theme.DisplayType) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.Display(display))
	}
}

func BackgroundColor(color string) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.BackgroundColor(color))
	}
}

func BorderColor(color string) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.BorderColor(color))
	}
}

func BorderStyle(style theme.BorderStyleType) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.BorderStyle(style))
	}
}

func BorderWidth(top int, right int, bottom int, left int) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.BorderWidth(top, right, bottom, left))
	}
}

func BorderRadius(borderRadius int) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.BorderRadius(borderRadius))
	}
}

func BoxShadow(shadow string) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.BoxShadow(shadow))
	}
}

func Class(class string) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.Class(class))
	}
}

func Flex(flex int) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.Flex(flex))
	}
}

func Height(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.Height(opts...))
	}
}

func ID(id string) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.ID(id))
	}
}

func Margin(opts ...breakpoint.BreakpointOptions[spacing.Spacing]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.Margin(opts...))
	}
}

func MaxWidth(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.MaxWidth(opts...))
	}
}

func Overflow(overflow theme.OverflowType) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.Overflow(overflow))
	}
}

func Padding(opts ...breakpoint.BreakpointOptions[spacing.Spacing]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.Padding(opts...))
	}
}

func Visible(opts ...breakpoint.BreakpointOptions[bool]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.Visible(opts...))
	}
}

func Width(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.Width(opts...))
	}
}

func WidthP(opts ...breakpoint.BreakpointOptions[float64]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.WidthP(opts...))
	}
}
