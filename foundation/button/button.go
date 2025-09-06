package button

import (
	basicbutton "github.com/gofred-io/gofred/basic/button"
	. "github.com/gofred-io/gofred/widget"
)

type FButton struct {
	*basicbutton.BButton
}

func Button(child Widget) *FButton {
	var (
		b = &FButton{
			BButton: basicbutton.Button(
				child,
			),
		}
	)

	b.Class("gf-button")
	return b
}
