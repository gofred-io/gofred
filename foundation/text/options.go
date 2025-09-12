package text

import (
	"github.com/gofred-io/gofred/basic/span"
	"github.com/gofred-io/gofred/theme"
	"github.com/gofred-io/gofred/theme/theme_data"
)

type Option func(text *text)

func setText(_text string) Option {
	return func(text *text) {
		text.opts = append(text.opts, span.SetText(_text))
	}
}

func Align(textAlign theme.TextAlignType) Option {
	return func(text *text) {
		text.opts = append(text.opts, span.TextAlign(textAlign))
	}
}

func Class(class string) Option {
	return func(text *text) {
		text.opts = append(text.opts, span.Class(class))
	}
}

func FontColor(color string) Option {
	return func(text *text) {
		text.opts = append(text.opts, span.FontColor(color))
	}
}

func FontFamily(fontFamily string) Option {
	return func(text *text) {
		text.opts = append(text.opts, span.FontFamily(fontFamily))
	}
}

func FontSize(fontSize int) Option {
	return func(text *text) {
		text.opts = append(text.opts, span.FontSize(fontSize))
	}
}

func FontWeight(fontWeight string) Option {
	return func(text *text) {
		text.opts = append(text.opts, span.FontWeight(fontWeight))
	}
}

func ID(id string) Option {
	return func(text *text) {
		text.opts = append(text.opts, span.ID(id))
	}
}

func LineHeight(lineHeight float64) Option {
	return func(text *text) {
		text.opts = append(text.opts, span.LineHeight(lineHeight))
	}
}

func TextStyle(textStyle theme_data.TextStyle) Option {
	return func(text *text) {
		text.opts = append(text.opts, span.TextStyle(textStyle))
	}
}

func UserSelect(userSelect theme.UserSelectType) Option {
	return func(text *text) {
		text.opts = append(text.opts, span.UserSelect(userSelect))
	}
}
