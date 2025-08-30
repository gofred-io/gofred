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

func ColumnCount(columnCount int) StyleOptions {
	return func(style *style.Style) {
		style.Display.ColumnCount = columnCount
	}
}

func Display(display style.Display) StyleOptions {
	return func(style *style.Style) {
		style.Display = &display
	}
}

func Height(height int) StyleOptions {
	return func(style *style.Style) {
		style.Size.Height = &height
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
	return func(style *style.Style) {
		style.Size.Width = &width
	}
}
