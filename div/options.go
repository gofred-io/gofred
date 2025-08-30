package div

import (
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/widget"
)

type Options func(div *Div)

func ID(id string) Options {
	return func(div *Div) {
		div.SetID(id)
	}
}

func Style(styleOptions ...StyleOptions) Options {
	return func(div *Div) {
		for _, styleOption := range styleOptions {
			styleOption(div.style)
		}
	}
}

func StyleFrom(style *style.Style) Options {
	return func(div *Div) {
		div.style = style
	}
}

func OnClick(onClick func(this widget.Widget)) Options {
	return func(div *Div) {
		div.onClick = onClick
	}
}
