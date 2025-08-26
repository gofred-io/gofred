package svg

import "github.com/gofred-io/gofred/widget"

type Svg struct {
	widget.Widget
	element widget.Widget
}

func New(children []widget.Widget, options ...Options) widget.Widget {
	svg := &Svg{
		Widget: widget.Context().CreateElement("span"),
	}

	svg.element = widget.Context().CreateElementNS("http://www.w3.org/2000/svg", "svg")

	for _, option := range options {
		option(svg)
	}

	svg.element.SetAttribute("xmlns", "http://www.w3.org/2000/svg")
	svg.element.SetAttribute("viewBox", "0 0 24 24")
	svg.element.SetAttribute("focusable", "false")
	svg.element.SetAttribute("aria-hidden", "true")
	svg.element.SetStyle("pointer-events: none; display: inherit;")

	svg.AppendChild(svg.element)

	for _, child := range children {
		svg.element.AppendChild(child)
	}

	return svg.Widget
}
