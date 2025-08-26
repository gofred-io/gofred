package svg

import (
	"strconv"

	"github.com/gofred-io/gofred/widget"
)

type Options func(svg *Svg)

func Children(children ...widget.Widget) Options {
	return func(svg *Svg) {
		svg.children = children
	}
}

func Height(height int) Options {
	return func(svg *Svg) {
		svg.element.SetAttribute("height", strconv.Itoa(height))
	}
}

func Width(width int) Options {
	return func(svg *Svg) {
		svg.element.SetAttribute("width", strconv.Itoa(width))
	}
}
