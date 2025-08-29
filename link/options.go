package link

import "github.com/gofred-io/gofred/widget"

type Options func(link *Link)

func OnClick(callback func(widget widget.Widget)) Options {
	return func(link *Link) {
		link.onClick = callback
	}
}
