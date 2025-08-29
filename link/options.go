package link

import "github.com/gofred-io/gofred/widget"

type Options func(link *Link)

func Href(href string) Options {
	return func(link *Link) {
		link.href = href
	}
}

func NewTab(newTab bool) Options {
	return func(link *Link) {
		link.newTab = newTab
	}
}

func OnClick(callback func(widget widget.Widget)) Options {
	return func(link *Link) {
		link.onClick = callback
	}
}
