package span

import (
	"github.com/gofred-io/gofred/widget"
)

type span struct {
	widget.BaseWidget
}

func New(child widget.BaseWidget, options ...Option) widget.BaseWidget {
	span := &span{
		BaseWidget: widget.New("span"),
	}

	for _, option := range options {
		option()(span.BaseWidget)
	}

	span.AppendChild(child.Widget)
	return span.BaseWidget
}
