package iconbutton

import "github.com/gofred-io/gofred/widget"

type Options func(iconButton *IconButton)

func OnClick(onClick func(this widget.Widget)) Options {
	return func(iconButton *IconButton) {
		iconButton.SetOnClick(onClick)
	}
}

func Style(styleOptions ...StyleOptions) Options {
	return func(iconButton *IconButton) {
		for _, styleOption := range styleOptions {
			styleOption(&iconButton.style)
		}
	}
}

func Tooltip(tooltip string) Options {
	return func(iconButton *IconButton) {
		iconButton.tooltip = tooltip
	}
}
