package iconbutton

import (
	"github.com/gofred-io/gofred/icon"
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/widget"
)

type IconButton struct {
	widget.Widget
	style   style.Style
	tooltip string
}

func New(iconData icondata.IconData, options ...Options) widget.Widget {
	iconButton := &IconButton{
		Widget: widget.Context().CreateElement("button"),
	}

	for _, option := range options {
		option(iconButton)
	}

	if iconButton.tooltip != "" {
		iconButton.SetAttribute("title", iconButton.tooltip)
	}

	iconButton.SetClass("gf-icon-button")
	iconButton.SetStyle(iconButton.style.String())

	iconElement := icon.New(iconData)
	iconButton.AppendChild(iconElement)

	return iconButton.Widget
}
