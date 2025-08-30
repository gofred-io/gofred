package container

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

func Flex(flex int) StyleOptions {
	return func(_style *style.Style) {
		if _style.Display == nil {
			_style.Display = &style.Display{}
		}

		_style.Display.Flex = flex
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

func HeightP(height float32) StyleOptions {
	return func(_style *style.Style) {
		if _style.Size == nil {
			_style.Size = &style.Size{}
		}

		_style.Size.HeightP = &height
	}
}

func Margin(margin style.Margin) StyleOptions {
	return func(style *style.Style) {
		style.Margin = &margin
	}
}

func MaxHeight(maxHeight int) StyleOptions {
	return func(_style *style.Style) {
		if _style.Size == nil {
			_style.Size = &style.Size{}
		}

		_style.Size.MaxHeight = &maxHeight
	}
}

func MaxWidth(maxWidth int) StyleOptions {
	return func(_style *style.Style) {
		if _style.Size == nil {
			_style.Size = &style.Size{}
		}

		_style.Size.MaxWidth = &maxWidth
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
