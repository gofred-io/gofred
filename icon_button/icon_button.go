package iconbutton

import (
	"github.com/gofred-io/gofred/icon"
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/widget"
)

type IconButton struct {
	widget.Widget
	style style.Style
}

func New(iconData icondata.IconData, options ...Options) widget.Widget {
	iconButton := &IconButton{
		Widget: widget.Context().CreateElement("button"),
	}

	for _, option := range options {
		option(iconButton)
	}

	iconButton.SetClass("gf-icon-button")
	iconButton.SetStyle(iconButton.style.String())

	iconElement := icon.New(iconData)
	iconButton.AppendChild(iconElement)

	return iconButton.Widget
}
