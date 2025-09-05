package pre

import (
	"github.com/gofred-io/gofred/widget"
)

type pre struct {
	*widget.BaseWidget
}

func New(child widget.Widget, opts ...Option) *pre {
	pre := &pre{
		BaseWidget: widget.New("pre"),
	}

	for _, option := range opts {
		option()(pre)
	}

	pre.AppendChild(child)
	return pre
}
