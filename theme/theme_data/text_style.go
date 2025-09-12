package theme_data

import (
	"strconv"

	"github.com/gofred-io/gofred/theme"
)

type TextStyleCollection struct {
	Error     TextStyle
	Footer    TextStyle
	Header    TextStyle
	Primary   TextStyle
	Secondary TextStyle
	Subheader TextStyle
	Tertiary  TextStyle
}

type TextStyle struct {
	BackgroundColor themeValue[string]
	Color           themeValue[string]
	FontSize        themeValue[int]
	FontWeight      themeValue[string]
	FontFamily      themeValue[string]
	LetterSpacing   themeValue[float64]
	LineHeight      themeValue[float64]
	Overflow        themeValue[theme.OverflowType]
	TextAlign       themeValue[theme.TextAlignType]
}

func (s TextStyle) Apply(widget Widget) {
	if s.BackgroundColor != nil {
		widget.UpdateStyleProperty("background-color", *s.BackgroundColor)
	}

	if s.Color != nil {
		widget.UpdateStyleProperty("color", *s.Color)
	}

	if s.FontSize != nil {
		widget.UpdateStyleProperty("font-size", strconv.Itoa(*s.FontSize))
	}

	if s.FontWeight != nil {
		widget.UpdateStyleProperty("font-weight", *s.FontWeight)
	}

	if s.FontFamily != nil {
		widget.UpdateStyleProperty("font-family", *s.FontFamily)
	}

	if s.LetterSpacing != nil {
		widget.UpdateStyleProperty("letter-spacing", strconv.FormatFloat(*s.LetterSpacing, 'f', -1, 64))
	}

	if s.LineHeight != nil {
		widget.UpdateStyleProperty("line-height", strconv.FormatFloat(*s.LineHeight, 'f', -1, 64))
	}

	if s.Overflow != nil {
		widget.UpdateStyleProperty("overflow", string(*s.Overflow))
	}

	if s.TextAlign != nil {
		widget.UpdateStyleProperty("text-align", string(*s.TextAlign))
	}
}
