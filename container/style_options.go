package container

import "github.com/gofred-io/gofred/style"

type StyleOptions func(style *style.Style)

func Margin(margin style.Margin) StyleOptions {
	return func(style *style.Style) {
		style.Margin = &margin
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
