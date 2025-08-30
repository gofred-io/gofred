package iconbutton

import (
	"github.com/gofred-io/gofred/style/breakpoint"
	"github.com/gofred-io/gofred/widget"
)

type Options func(iconButton *IconButton)

func Class(class string) Options {
	return func(iconButton *IconButton) {
		iconButton.AddClass(class)
	}
}

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

func Visible(options ...breakpoint.BreakpointOptions[bool]) Options {
	return func(iconButton *IconButton) {
		if iconButton.visible == nil {
			iconButton.visible = &breakpoint.BreakpointValue[bool]{}
		}

		for _, visibleOption := range options {
			visibleOption(iconButton.visible)
		}
	}
}
