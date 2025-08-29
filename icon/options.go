package icon

type Options func(icon *Icon)

func Style(styleOptions ...StyleOptions) Options {
	return func(icon *Icon) {
		for _, styleOption := range styleOptions {
			styleOption(&icon.style)
		}
	}
}
