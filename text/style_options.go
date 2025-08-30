package text

import (
	"github.com/gofred-io/gofred/style"
)

type StyleOptions func(style *style.Style)

func Font(fontOptions ...FontOptions) StyleOptions {
	return func(_style *style.Style) {
		font := &style.Font{}
		for _, fontOption := range fontOptions {
			fontOption(font)
		}
		_style.Font = font
	}
}

func LineHeight(lineHeight float64) StyleOptions {
	return func(_style *style.Style) {
		_style.LineHeight = &lineHeight
	}
}

func UserSelect(userSelect style.UserSelectType) StyleOptions {
	return func(_style *style.Style) {
		if _style.Display == nil {
			_style.Display = &style.Display{}
		}
		_style.Display.UserSelect = userSelect
	}
}
