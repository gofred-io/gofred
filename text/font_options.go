package text

import "github.com/gofred-io/gofred/style"

type FontOptions func(font *style.Font)

func Color(color string) FontOptions {
	return func(font *style.Font) {
		font.Color = color
	}
}

func Family(family string) FontOptions {
	return func(font *style.Font) {
		font.Family = family
	}
}

func Size(size int) FontOptions {
	return func(font *style.Font) {
		font.Size = size
	}
}

func Weight(weight string) FontOptions {
	return func(font *style.Font) {
		font.Weight = weight
	}
}
