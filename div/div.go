package div

import (
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type Div struct {
	widget.BaseWidget
}

func New(children []widget.BaseWidget, opts ...options.Options) widget.BaseWidget {
	div := &Div{
		BaseWidget: widget.New("div"),
	}

	for _, option := range opts {
		option(div.BaseWidget)
	}

	for _, child := range children {
		div.AppendChild(child.Widget)
	}

	return div.BaseWidget
}
