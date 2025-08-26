package div

import "github.com/gofred-io/gofred/style"

type StyleOptions func(style *style.Style)

func Background(background style.Background) StyleOptions {
	return func(style *style.Style) {
		style.Background = &background
	}
}

func Border(border style.Border) StyleOptions {
	return func(style *style.Style) {
		style.Border = &border
	}
}

func Display(display style.Display) StyleOptions {
	return func(style *style.Style) {
		style.Display = &display
	}
}

func Padding(padding style.Padding) StyleOptions {
	return func(style *style.Style) {
		style.Padding = &padding
	}
}

func Size(width, height int) StyleOptions {
	return func(_style *style.Style) {
		_style.Size = &style.Size{
			Width:  width,
			Height: height,
		}
	}
}
