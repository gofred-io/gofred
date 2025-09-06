package button

import (
	. "github.com/gofred-io/gofred/widget"
)

type BButton struct {
	*BaseWidget
}

func Button(child Widget) *BButton {
	button := &BButton{
		BaseWidget: New("button"),
	}

	button.AppendChild(child)
	return button
}
