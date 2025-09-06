package span

import (
	. "github.com/gofred-io/gofred/widget"
)

type BSpan struct {
	*BaseWidget
}

func Span(child Widget) *BSpan {
	span := &BSpan{
		BaseWidget: New("span"),
	}

	if child != nil {
		span.AppendChild(child)
	}

	return span
}
