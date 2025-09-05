package span

import (
	"github.com/gofred-io/gofred/widget"
)

type span struct {
	*widget.BaseWidget
}

func New(child widget.Widget, options ...Option) *span {
	span := &span{
		BaseWidget: widget.New("span"),
	}

	for _, option := range options {
		option()(span)
	}

	if child != nil {
		span.AppendChild(child)
	}

	return span
}
