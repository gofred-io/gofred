package svg

import (
	"github.com/gofred-io/gofred/application"
)

type svg struct {
	application.BaseWidget
}

func New(children []application.BaseWidget, opts ...Option) application.BaseWidget {
	svg := &svg{
		BaseWidget: application.NewNS("http://www.w3.org/2000/svg", "svg"),
	}

	for _, option := range opts {
		option()(svg.BaseWidget)
	}

	svg.SetAttribute("xmlns", "http://www.w3.org/2000/svg")
	svg.SetAttribute("viewBox", "0 0 24 24")

	for _, child := range children {
		svg.AppendChild(child.Widget)
	}

	return svg.BaseWidget
}
