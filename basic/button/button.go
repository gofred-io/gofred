package button

import (
	"github.com/gofred-io/gofred/widget"
)

type button struct {
	*widget.BaseWidget
}

func New(child widget.Widget, options ...Option) *button {
	button := &button{
		BaseWidget: widget.New("button"),
	}

	for _, opt := range options {
		opt()(button)
	}

	button.AppendChild(child)
	return button
}
