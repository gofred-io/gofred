package iconbutton

import (
	"github.com/gofred-io/gofred/basic/button"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/icon"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/theme"
	"github.com/gofred-io/gofred/theme/theme_style"
)

type Option func(iconButton *iconButton)

func BorderRadius(borderRadius int) Option {
	return func(iconButton *iconButton) {
		iconButton.opts = append(iconButton.opts, button.BorderRadius(borderRadius))
	}
}

func ButtonStyle(buttonStyle theme_style.ButtonStyle) Option {
	return func(iconButton *iconButton) {
		iconButton.opts = append(iconButton.opts, button.ButtonStyle(buttonStyle))
	}
}

func Class(class string) Option {
	return func(iconButton *iconButton) {
		iconButton.opts = append(iconButton.opts, button.Class(class))
	}
}

func Fill(fill string) Option {
	return func(iconButton *iconButton) {
		iconButton.iconOpts = append(iconButton.iconOpts, icon.Fill(fill))
	}
}

func Height(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(iconButton *iconButton) {
		iconButton.opts = append(iconButton.opts, button.Height(opts...))
	}
}

func ID(id string) Option {
	return func(iconButton *iconButton) {
		iconButton.opts = append(iconButton.opts, button.ID(id))
	}
}

func Label(label string) Option {
	return func(iconButton *iconButton) {
		iconButton.opts = append(iconButton.opts, button.Label(label))
	}
}

func MaxWidth(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(iconButton *iconButton) {
		iconButton.opts = append(iconButton.opts, button.MaxWidth(opts...))
	}
}

func OnClick(handler options.OnClickHandler) Option {
	return func(iconButton *iconButton) {
		iconButton.opts = append(iconButton.opts, button.OnClick(handler))
	}
}

func Tooltip(tooltip string) Option {
	return func(iconButton *iconButton) {
		iconButton.opts = append(iconButton.opts, button.Tooltip(tooltip))
	}
}

func UserSelect(userSelect theme.UserSelectType) Option {
	return func(iconButton *iconButton) {
		iconButton.iconOpts = append(iconButton.iconOpts, icon.UserSelect(userSelect))
	}
}
func Visible(opts ...breakpoint.BreakpointOptions[bool]) Option {
	return func(iconButton *iconButton) {
		iconButton.opts = append(iconButton.opts, button.Visible(opts...))
	}
}

func Width(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(iconButton *iconButton) {
		iconButton.opts = append(iconButton.opts, button.Width(opts...))
	}
}

func WidthP(opts ...breakpoint.BreakpointOptions[float64]) Option {
	return func(iconButton *iconButton) {
		iconButton.opts = append(iconButton.opts, button.WidthP(opts...))
	}
}
