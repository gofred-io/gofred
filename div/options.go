package div

import (
	"github.com/man.go/mango/widget"
)

type Options func(div *Div)

func Children(children ...widget.Widget) Options {
	return func(div *Div) {
		div.children = children
	}
}

func Style(styleOptions ...StyleOptions) Options {
	return func(div *Div) {
		for _, styleOption := range styleOptions {
			styleOption(&div.style)
		}
	}
}

func OnClick(onClick func(this widget.Widget)) Options {
	return func(div *Div) {
		div.onClick = onClick
	}
}
