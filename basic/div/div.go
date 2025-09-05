package div

import (
	"github.com/gofred-io/gofred/widget"
)

type div struct {
	*widget.BaseWidget
}

func New(children []widget.Widget, opts ...Option) *div {
	div := &div{
		BaseWidget: widget.New("div"),
	}

	for _, option := range opts {
		option()(div)
	}

	for _, child := range children {
		div.AppendChild(child)
	}

	return div
}
