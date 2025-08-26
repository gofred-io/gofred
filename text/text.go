package text

import (
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/widget"
)

type text struct {
	widget.Widget
	style style.Style
}

func New(options ...Options) widget.Widget {
	text := &text{
		Widget: widget.Context().CreateElement("span"),
	}

	for _, option := range options {
		option(text)
	}

	text.SetStyle(text.style.String())

	return text.Widget
}
