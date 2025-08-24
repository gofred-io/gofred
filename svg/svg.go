package svg

import "github.com/man.go/mango/widget"

type Svg struct {
	widget.Widget
	element  widget.Widget
	children []widget.Widget
}

func New(ctx *widget.WidgetContext, options ...Options) widget.Widget {
	svg := &Svg{
		Widget: ctx.CreateElement("span"),
	}

	svg.element = ctx.CreateElementNS("http://www.w3.org/2000/svg", "svg")

	for _, option := range options {
		option(svg)
	}

	svg.element.SetAttribute("xmlns", "http://www.w3.org/2000/svg")
	svg.element.SetAttribute("viewBox", "0 0 24 24")
	svg.element.SetAttribute("focusable", "false")
	svg.element.SetAttribute("aria-hidden", "true")
	svg.element.SetStyle("pointer-events: none; display: inherit;")

	svg.AppendChild(svg.element)

	for _, child := range svg.children {
		svg.element.AppendChild(child)
	}

	return svg.Widget
}
