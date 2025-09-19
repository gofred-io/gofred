package button

import (
	"github.com/gofred-io/gofred/application"
	basicbutton "github.com/gofred-io/gofred/basic/button"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/theme/theme_style"
)

type button struct {
	opts []basicbutton.Option
}

func New(child application.BaseWidget, opts ...Option) application.BaseWidget {
	var (
		b = &button{}
	)

	opts = mergeWithDefaultOpts(opts)
	for _, option := range opts {
		option(b)
	}

	return basicbutton.New(
		child,
		b.opts...,
	)
}

func mergeWithDefaultOpts(opts []Option) []Option {
	return append(
		defaultOpts(),
		opts...,
	)
}

func defaultOpts() []Option {
	return []Option{
		ButtonStyle(defaultThemeData()),
		Class("gf-button"),
	}
}

func defaultThemeData() theme_style.ButtonStyle {
	themeHook, _ := hooks.UseTheme()
	themeData := themeHook.ThemeData()
	return themeData.ButtonTheme.ButtonStyle.Primary
}
