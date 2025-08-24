package div

import (
	"github.com/man.go/mango/style"
	"github.com/man.go/mango/widget"
)

type div struct {
	widget.Widget
	children []widget.Widget
	style    style.Style

	// events
	onClick func(this widget.Widget)
}

func New(ctx *widget.WidgetContext, options ...Options) widget.Widget {
	div := &div{
		Widget: ctx.CreateElement("div"),
	}

	for _, option := range options {
		option(div)
	}

	div.SetStyle(div.style.String())
	div.SetOnClick(div.onClick)

	for _, child := range div.children {
		div.AppendChild(child)
	}

	return div.Widget
}
