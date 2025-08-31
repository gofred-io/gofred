package text

import (
	"github.com/gofred-io/gofred/widget"
)

type text struct {
	widget.BaseWidget
}

func New(innerText string, options ...Option) widget.BaseWidget {
	_text := &text{
		BaseWidget: widget.New("span"),
	}

	for _, option := range options {
		option()(_text.BaseWidget)
	}

	_text.SetText(innerText)

	return _text.BaseWidget
}
