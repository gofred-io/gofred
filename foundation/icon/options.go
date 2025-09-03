package icon

import (
	"github.com/gofred-io/gofred/basic/svg"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/options"
)

type Option func(icon *icon)

func Class(class string) Option {
	return func(icon *icon) {
		icon.opts = append(icon.opts, svg.Class(class))
	}
}

func Fill(fill string) Option {
	return func(icon *icon) {
		icon.opts = append(icon.opts, svg.Fill(fill))
	}
}

func Height(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(icon *icon) {
		icon.opts = append(icon.opts, svg.Height(opts...))
	}
}

func ID(id string) Option {
	return func(icon *icon) {
		icon.opts = append(icon.opts, svg.ID(id))
	}
}

func MaxWidth(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(icon *icon) {
		icon.opts = append(icon.opts, svg.MaxWidth(opts...))
	}
}

func UserSelect(userSelect options.UserSelectType) Option {
	return func(icon *icon) {
		icon.opts = append(icon.opts, svg.UserSelect(userSelect))
	}
}

func Width(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(icon *icon) {
		icon.opts = append(icon.opts, svg.Width(opts...))
	}
}

func WidthP(opts ...breakpoint.BreakpointOptions[float64]) Option {
	return func(icon *icon) {
		icon.opts = append(icon.opts, svg.WidthP(opts...))
	}
}
