package text

import (
	"github.com/man.go/mango/style"
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
