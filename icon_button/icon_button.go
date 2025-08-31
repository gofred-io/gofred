package iconbutton

import (
	"github.com/gofred-io/gofred/icon"
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type IconButton struct {
	widget.BaseWidget
}

func New(iconData icondata.IconData, opts ...options.Options) widget.BaseWidget {
	iconButton := &IconButton{
		BaseWidget: widget.New("button"),
	}

	iconButton.AddClass("gf-icon-button")

	for _, option := range opts {
		option(iconButton.BaseWidget)
	}

	iconElement := icon.New(iconData)
	iconButton.AppendChild(iconElement.Widget)

	return iconButton.BaseWidget
}
