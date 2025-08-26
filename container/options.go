package container

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
