package pre

import (
	"github.com/gofred-io/gofred/widget"
)

type pre struct {
	widget.BaseWidget
}

func New(child widget.BaseWidget, opts ...Option) widget.BaseWidget {
	pre := &pre{
		BaseWidget: widget.New("pre"),
	}

	for _, option := range opts {
		option()(pre.BaseWidget)
	}

	pre.AppendChild(child.Widget)
	return pre.BaseWidget
}
