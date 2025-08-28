package image

type Options func(image *Image)

func Alt(alt string) Options {
	return func(image *Image) {
		image.Set("alt", alt)
	}
}

func Height(height int) Options {
	return func(image *Image) {
		image.Set("height", height)
	}
}

func Width(width int) Options {
	return func(image *Image) {
		image.Set("width", width)
	}
}
