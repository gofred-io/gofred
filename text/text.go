package text

import (
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/widget"
)

type text struct {
	widget.Widget
	style style.Style
}

func New(innerText string, options ...Options) widget.Widget {
	_text := &text{
		Widget: widget.Context().CreateElement("span"),
	}

	for _, option := range options {
		option(_text)
	}

	_text.SetText(innerText)
	_text.SetStyle(_text.style.String())

	return _text.Widget
}
