package button

import "github.com/gofred-io/gofred/widget"

type Options func(button *Button)

func OnClick(onClick func(this widget.Widget)) Options {
	return func(button *Button) {
		button.SetOnClick(onClick)
	}
}

func Style(styleOptions ...StyleOptions) Options {
	return func(button *Button) {
		for _, styleOption := range styleOptions {
			styleOption(&button.style)
		}
	}
}

func Tooltip(tooltip string) Options {
	return func(button *Button) {
		button.tooltip = tooltip
	}
}
