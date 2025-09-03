package button

import (
	basicbutton "github.com/gofred-io/gofred/basic/button"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/options"
)

type Option func(button *button)

func BackgroundColor(backgroundColor string) Option {
	return func(button *button) {
		button.opts = append(button.opts, basicbutton.BackgroundColor(backgroundColor))
	}
}

func BorderRadius(borderRadius int) Option {
	return func(button *button) {
		button.opts = append(button.opts, basicbutton.BorderRadius(borderRadius))
	}
}

func Class(class string) Option {
	return func(button *button) {
		button.opts = append(button.opts, basicbutton.Class(class))
	}
}

func Height(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(button *button) {
		button.opts = append(button.opts, basicbutton.Height(opts...))
	}
}

func ID(id string) Option {
	return func(button *button) {
		button.opts = append(button.opts, basicbutton.ID(id))
	}
}

func MaxWidth(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(button *button) {
		button.opts = append(button.opts, basicbutton.MaxWidth(opts...))
	}
}

func OnClick(handler options.OnClickHandler) Option {
	return func(button *button) {
		button.opts = append(button.opts, basicbutton.OnClick(handler))
	}
}

func Tooltip(tooltip string) Option {
	return func(button *button) {
		button.opts = append(button.opts, basicbutton.Tooltip(tooltip))
	}
}

func Visible(opts ...breakpoint.BreakpointOptions[bool]) Option {
	return func(button *button) {
		button.opts = append(button.opts, basicbutton.Visible(opts...))
	}
}

func Width(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(button *button) {
		button.opts = append(button.opts, basicbutton.Width(opts...))
	}
}

func WidthP(opts ...breakpoint.BreakpointOptions[float64]) Option {
	return func(button *button) {
		button.opts = append(button.opts, basicbutton.WidthP(opts...))
	}
}
