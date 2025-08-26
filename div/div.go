package div

import (
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/widget"
)

type Div struct {
	widget.Widget
	style style.Style

	// events
	onClick func(this widget.Widget)
}

func New(children []widget.Widget, options ...Options) widget.Widget {
	div := &Div{
		Widget: widget.Context().CreateElement("div"),
	}

	for _, option := range options {
		option(div)
	}

	div.SetStyle(div.style.String())
	div.SetOnClick(div.onClick)

	for _, child := range children {
		div.AppendChild(child)
	}

	return div.Widget
}
