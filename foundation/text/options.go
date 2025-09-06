package text

import (
	. "github.com/gofred-io/gofred/options"
)

func (text *FText) setText(_text string) *FText {
	text.BSpan.SetText(_text)
	return text
}

func (text *FText) Align(textAlign TextAlignType) *FText {
	text.BSpan.TextAlign(textAlign)
	return text
}

func (text *FText) Class(class string) *FText {
	text.BSpan.Class(class)
	return text
}

func (text *FText) FontColor(color string) *FText {
	text.BSpan.FontColor(color)
	return text
}

func (text *FText) FontFamily(fontFamily string) *FText {
	text.BSpan.FontFamily(fontFamily)
	return text
}

func (text *FText) FontSize(fontSize int) *FText {
	text.BSpan.FontSize(fontSize)
	return text
}

func (text *FText) FontWeight(fontWeight string) *FText {
	text.BSpan.FontWeight(fontWeight)
	return text
}

func (text *FText) ID(id string) *FText {
	text.BSpan.ID(id)
	return text
}

func (text *FText) LineHeight(lineHeight float64) *FText {
	text.BSpan.LineHeight(lineHeight)
	return text
}

func (text *FText) UserSelect(userSelect UserSelectType) *FText {
	text.BSpan.UserSelect(userSelect)
	return text
}
