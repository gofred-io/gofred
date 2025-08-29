package link

import "github.com/gofred-io/gofred/widget"

type Link struct {
	widget.Widget

	onClick func(widget widget.Widget)
}

func New(child widget.Widget, options ...Options) widget.Widget {
	link := &Link{
		Widget: widget.Context().CreateElement("span"),
	}

	link.SetStyle("cursor: pointer;")
	defaultColor := link.Get("style").Get("color").String()
	link.SetAttribute("onmouseover", "this.style.color = '#0000EE';")
	link.SetAttribute("onmouseout", "this.style.color = '"+defaultColor+"';")

	for _, option := range options {
		option(link)
	}

	link.SetOnClick(func(widget widget.Widget) {
		if link.onClick != nil {
			link.onClick(widget)
		}
	})

	link.AppendChild(child)
	return link.Widget
}
