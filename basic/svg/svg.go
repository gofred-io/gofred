package svg

import (
	. "github.com/gofred-io/gofred/widget"
)

type BSvg struct {
	*BaseWidget
}

func Svg(children []Widget) *BSvg {
	svg := &BSvg{
		BaseWidget: NewNS("http://www.w3.org/2000/svg", "svg"),
	}

	svg.SetAttribute("xmlns", "http://www.w3.org/2000/svg")
	svg.SetAttribute("viewBox", "0 0 24 24")

	for _, child := range children {
		svg.AppendChild(child)
	}

	return svg
}
