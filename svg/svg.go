package svg

import (
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type Svg struct {
	widget.BaseWidget
}

func New(children []widget.BaseWidget, options ...options.Options) widget.BaseWidget {
	svg := &Svg{
		BaseWidget: widget.NewNS("http://www.w3.org/2000/svg", "svg"),
	}

	for _, option := range options {
		option(svg.BaseWidget)
	}

	svg.SetAttribute("xmlns", "http://www.w3.org/2000/svg")
	svg.SetAttribute("viewBox", "0 0 24 24")

	for _, child := range children {
		svg.AppendChild(child.Widget)
	}

	return svg.BaseWidget
}
