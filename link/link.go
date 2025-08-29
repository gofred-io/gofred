package link

import "github.com/gofred-io/gofred/widget"

type Link struct {
	widget.Widget
	href    string
	newTab  bool
	tooltip string

	onClick func(widget widget.Widget)
}

func New(child widget.Widget, options ...Options) widget.Widget {
	link := &Link{
		Widget: widget.Context().CreateElement("a"),
	}

	link.SetClass("gf-link")

	for _, option := range options {
		option(link)
	}

	link.SetAttribute("href", link.href)
	if link.newTab {
		link.SetAttribute("target", "_blank")
	}

	if link.tooltip != "" {
		link.SetAttribute("title", link.tooltip)
	}

	link.SetOnClick(func(widget widget.Widget) {
		if link.onClick != nil {
			link.onClick(widget)
		}
	})

	link.AppendChild(child)
	return link.Widget
}
