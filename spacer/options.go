package spacer

type Options func(spacer *Spacer)

func Height(height int) Options {
	return func(spacer *Spacer) {
		spacer.height = height
	}
}

func Width(width int) Options {
	return func(spacer *Spacer) {
		spacer.width = width
	}
}
