package div

import (
	"github.com/man.go/mango/widget"
)

type Options func(div *div)

func Children(children ...widget.Widget) Options {
	return func(div *div) {
		div.children = children
	}
}

func ID(id string) Options {
	return func(div *div) {
		div.SetID(id)
	}
}

func Style(styleOptions ...StyleOptions) Options {
	return func(div *div) {
		for _, styleOption := range styleOptions {
			styleOption(&div.style)
		}
	}
}

func OnClick(onClick func(this widget.Widget)) Options {
	return func(div *div) {
		div.onClick = onClick
	}
}
