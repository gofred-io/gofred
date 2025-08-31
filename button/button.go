package button

import (
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type Button struct {
	widget.BaseWidget
}

func New(child widget.BaseWidget, options ...options.Options) widget.BaseWidget {
	button := &Button{
		BaseWidget: widget.New("button"),
	}

	for _, option := range options {
		option(button.BaseWidget)
	}

	button.AddClass("gf-button")
	button.AppendChild(child.Widget)

	return button.BaseWidget
}
