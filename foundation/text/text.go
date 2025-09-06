package text

import (
	. "github.com/gofred-io/gofred/basic/span"
)

type FText struct {
	*BSpan
}

func Text(innerText string) *FText {
	text := &FText{
		BSpan: Span(nil),
	}

	text.setText(innerText)
	return text
}
