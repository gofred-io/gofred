package div

import (
	"github.com/google/uuid"
	"github.com/man.go/mango/style"
	"github.com/man.go/mango/widget"
)

type Div struct {
	widget.Widget
	children []widget.Widget
	style    style.Style

	// events
	onClick func(this widget.Widget)
}

func New(ctx *widget.WidgetContext, options ...Options) widget.Widget {
	div := &Div{
		Widget: ctx.CreateElement("div"),
	}

	for _, option := range options {
		option(div)
	}

	div.SetID(uuid.New().String())
	div.SetStyle(div.style.String())
	div.SetOnClick(div.onClick)

	for _, child := range div.children {
		div.AppendChild(child)
	}

	return div.Widget
}
