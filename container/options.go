package container

import "github.com/gofred-io/gofred/style"

type Options func(container *Container)

func ID(id string) Options {
	return func(container *Container) {
		container.SetID(id)
	}
}

func Margin(margin style.Margin) Options {
	return func(container *Container) {
		container.style.Margin = &margin
	}
}

func Padding(padding style.Padding) Options {
	return func(container *Container) {
		container.style.Padding = &padding
	}
}
