package svg

import "strconv"

type StyleOptions func(svg *Svg)

func Height(height int) StyleOptions {
	return func(svg *Svg) {
		svg.element.SetAttribute("height", strconv.Itoa(height))
	}
}

func Width(width int) StyleOptions {
	return func(svg *Svg) {
		svg.element.SetAttribute("width", strconv.Itoa(width))
	}
}
