package container

import (
	"github.com/gofred-io/gofred/style/breakpoint"
	"github.com/gofred-io/gofred/widget"
)

type Options func(container *Container)

func ID(id string) Options {
	return func(container *Container) {
		container.SetID(id)
	}
}

func Style(styleOptions ...StyleOptions) Options {
	return func(container *Container) {
		for _, styleOption := range styleOptions {
			styleOption(&container.style)
		}
	}
}

func OnClick(callback func(widget widget.Widget)) Options {
	return func(container *Container) {
		container.onClick = callback
	}
}

func Visible(options ...breakpoint.BreakpointOptions[bool]) Options {
	return func(container *Container) {
		if container.visible == nil {
			container.visible = &breakpoint.BreakpointValue[bool]{}
		}

		for _, visibleOption := range options {
			visibleOption(container.visible)
		}
	}
}

func WidthP(options ...breakpoint.BreakpointOptions[float64]) Options {
	return func(container *Container) {
		if container.width == nil {
			container.width = &breakpoint.BreakpointValue[float64]{}
		}

		for _, widthOption := range options {
			widthOption(container.width)
		}
	}
}
