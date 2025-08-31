package container

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/div"
	"github.com/gofred-io/gofred/options"
)

type Option func(container *container)

func display(display options.DisplayType) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.Display(display))
	}
}

func BackgroundColor(color string) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.BackgroundColor(color))
	}
}

func BorderRadius(borderRadius int) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.BorderRadius(borderRadius))
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

func Margin(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.Margin(opts...))
	}
}

func MarginB(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.MarginB(opts...))
	}
}

func MarginH(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.MarginH(opts...))
	}
}

func MarginL(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.MarginL(opts...))
	}
}

func MarginR(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.MarginR(opts...))
	}
}

func MarginT(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.MarginT(opts...))
	}
}

func MarginV(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.MarginV(opts...))
	}
}

func MaxWidth(maxWidth int) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.MaxWidth(maxWidth))
	}
}

func Padding(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.Padding(opts...))
	}
}

func PaddingB(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.PaddingB(opts...))
	}
}

func PaddingH(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.PaddingH(opts...))
	}
}

func PaddingL(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.PaddingL(opts...))
	}
}

func PaddingR(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.PaddingR(opts...))
	}
}

func PaddingT(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.PaddingT(opts...))
	}
}

func PaddingV(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(container *container) {
		container.opts = append(container.opts, div.PaddingV(opts...))
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
