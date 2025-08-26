package svg

import (
	"strconv"
)

type Options func(svg *Svg)

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
