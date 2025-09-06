package div

import (
	. "github.com/gofred-io/gofred/widget"
)

type BDiv struct {
	*BaseWidget
}

func Div(children []Widget) *BDiv {
	div := &BDiv{
		BaseWidget: New("div"),
	}

	for _, child := range children {
		div.AppendChild(child)
	}

	return div
}
