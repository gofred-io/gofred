package link

import (
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type Link struct {
	widget.BaseWidget
}

func New(child widget.BaseWidget, opts ...options.Options) widget.BaseWidget {
	link := &Link{
		BaseWidget: widget.NewWithOptions("a", map[string]any{
			"is": "pushstate-anchor",
		}),
	}

	link.AddClass("gf-link")

	for _, option := range opts {
		option(link.BaseWidget)
	}

	link.AppendChild(child.Widget)
	return link.BaseWidget
}
