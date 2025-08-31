package button

import (
	"github.com/gofred-io/gofred/widget"
)

type button struct {
	widget.BaseWidget
}

func New(child widget.BaseWidget, options ...Option) widget.BaseWidget {
	button := &button{
		BaseWidget: widget.New("button"),
	}

	for _, opt := range options {
		opt()(button.BaseWidget)
	}

	button.AddClass("gf-button")
	button.AppendChild(child.Widget)

	return button.BaseWidget
}
