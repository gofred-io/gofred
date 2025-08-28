package iconbutton

import "github.com/gofred-io/gofred/widget"

type Options func(iconButton *IconButton)

func OnClick(onClick func(this widget.Widget)) Options {
	return func(iconButton *IconButton) {
		iconButton.SetOnClick(onClick)
	}
}
