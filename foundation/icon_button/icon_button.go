package iconbutton

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/basic/button"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/theme/theme_style"
)

type iconButton struct {
	opts     []button.Option
	iconOpts []icon.Option
}

func New(iconData icondata.IconData, opts ...Option) application.BaseWidget {
	iconButton := &iconButton{}

	opts = mergeWithDefaultOpts(opts)
	for _, option := range opts {
		option(iconButton)
	}

	iconElement := icon.New(iconData, iconButton.iconOpts...)
	return button.New(iconElement, iconButton.opts...)
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
		Class("gf-icon-button"),
	}
}

func defaultThemeData() theme_style.ButtonStyle {
	themeHook, _ := hooks.UseTheme()
	themeData := themeHook.ThemeData()
	return themeData.ButtonTheme.IconButtonStyle.Primary
}
