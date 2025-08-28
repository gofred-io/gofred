package container

import "github.com/gofred-io/gofred/style"

type StyleOptions func(style *style.Style)

func Background(background style.Background) StyleOptions {
	return func(style *style.Style) {
		style.Background = &background
	}
}

func Height(height int) StyleOptions {
	return func(_style *style.Style) {
		if _style.Size == nil {
			_style.Size = &style.Size{}
		}

		_style.Size.Height = &height
	}
}

func Margin(margin style.Margin) StyleOptions {
	return func(style *style.Style) {
		style.Margin = &margin
	}
}

func Padding(padding style.Padding) StyleOptions {
	return func(style *style.Style) {
		style.Padding = &padding
	}
}

func Width(width int) StyleOptions {
	return func(_style *style.Style) {
		if _style.Size == nil {
			_style.Size = &style.Size{}
		}

		_style.Size.Width = &width
	}
}
