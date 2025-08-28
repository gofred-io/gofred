package container

import (
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
