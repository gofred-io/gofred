package svg

import (
	"github.com/gofred-io/gofred/widget"
)

type svg struct {
	*widget.BaseWidget
}

func New(children []widget.Widget, opts ...Option) *svg {
	svg := &svg{
		BaseWidget: widget.NewNS("http://www.w3.org/2000/svg", "svg"),
	}

	for _, option := range opts {
		option()(svg)
	}

	svg.SetAttribute("xmlns", "http://www.w3.org/2000/svg")
	svg.SetAttribute("viewBox", "0 0 24 24")

	for _, child := range children {
		svg.AppendChild(child)
	}

	return svg
}
