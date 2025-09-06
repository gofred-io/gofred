package iconbutton

import (
	. "github.com/gofred-io/gofred/basic/button"
	. "github.com/gofred-io/gofred/foundation/icon"
	. "github.com/gofred-io/gofred/foundation/icon/icon_data"
)

type FIconButton struct {
	*BButton
	*FIcon
}

func IconButton(iconData IconData) *FIconButton {
	iconButton := &FIconButton{
		FIcon: Icon(iconData),
	}

	iconButton.BButton = Button(
		iconButton.FIcon,
	)

	iconButton.Class("gf-icon-button")
	return iconButton
}
