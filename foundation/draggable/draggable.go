package draggable

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/theme"
	"github.com/gofred-io/gofred/theme/theme_style"
)

type draggable struct {
	opts []div.Option
}

func New(child application.BaseWidget, opts ...Option) application.BaseWidget {
	var (
		c = &draggable{}
	)

	opts = mergeWithDefaultOpts(opts)
	for _, option := range opts {
		option(c)
	}

	return div.New(
		[]application.BaseWidget{child},
		c.opts...,
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
		containerStyle(defaultThemeData()),
		display(theme.DisplayTypeFlex),
		setDraggable(true),
	}
}

func defaultThemeData() theme_style.ContainerStyle {
	themeHook, _ := hooks.UseTheme()
	themeData := themeHook.ThemeData()
	return themeData.BoxTheme.ContainerStyle.Primary
}
