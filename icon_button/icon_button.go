package iconbutton

import (
	"github.com/gofred-io/gofred/icon"
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/widget"
)

type IconButton struct {
	widget.Widget
}

func New(iconData icondata.IconData, options ...Options) widget.Widget {
	iconButton := &IconButton{
		Widget: widget.Context().CreateElement("button"),
	}

	iconButton.SetStyle("background-color: transparent; border: none; cursor: pointer; border-radius: 50%; padding: 8px; margin: 0;")
	iconButton.SetAttribute("onmouseover", "this.style.backgroundColor = '#F0F0F0';")
	iconButton.SetAttribute("onmouseout", "this.style.backgroundColor = 'transparent';")
	iconButton.SetAttribute("onmousedown", "this.style.backgroundColor = '#E8E8E8';")
	iconButton.SetAttribute("onmouseup", "this.style.backgroundColor = '#F0F0F0';")

	for _, option := range options {
		option(iconButton)
	}

	iconElement := icon.New(iconData)
	iconButton.AppendChild(iconElement)

	return iconButton.Widget
}
