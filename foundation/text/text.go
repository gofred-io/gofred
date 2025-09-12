package text

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/basic/span"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/theme/theme_data"
)

type text struct {
	opts []span.Option
}

func New(innerText string, opts ...Option) application.BaseWidget {
	text := &text{}

	opts = mergeWithDefaultOpts(opts)
	opts = append(opts, setText(innerText))

	for _, option := range opts {
		option(text)
	}

	return span.New(application.Nil, text.opts...)
}

func mergeWithDefaultOpts(opts []Option) []Option {
	return append(
		defaultOpts(),
		opts...,
	)
}

func defaultOpts() []Option {
	return []Option{
		TextStyle(defaultThemeData()),
	}
}

func defaultThemeData() theme_data.TextStyle {
	themeHook, _ := hooks.UseTheme()
	themeData := themeHook.ThemeData()
	return themeData.TextTheme().TextStyle().Primary
}
