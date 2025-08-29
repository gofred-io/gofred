package button

import (
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/widget"
)

type Button struct {
	widget.Widget
	style   style.Style
	tooltip string
}

func New(child widget.Widget, options ...Options) widget.Widget {
	button := &Button{
		Widget: widget.Context().CreateElement("button"),
	}

	for _, option := range options {
		option(button)
	}

	if button.tooltip != "" {
		button.SetAttribute("title", button.tooltip)
	}

	button.SetClass("gf-button")
	button.SetStyle(button.style.String())
	button.AppendChild(child)

	return button.Widget
}
