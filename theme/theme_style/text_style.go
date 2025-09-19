package theme_style

import (
	"github.com/gofred-io/gofred/theme"
	"github.com/gofred-io/gofred/theme/color"
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
	BackgroundColor themeValue[color.Color]
	Color           themeValue[color.Color]
	FontSize        themeValue[int]
	FontWeight      themeValue[string]
	FontFamily      themeValue[string]
	LetterSpacing   themeValue[float64]
	LineHeight      themeValue[float64]
	Overflow        themeValue[theme.OverflowType]
	TextAlign       themeValue[theme.TextAlignType]
}
