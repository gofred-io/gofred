package link

import (
	"github.com/gofred-io/gofred/widget"
)

type Link struct {
	widget.BaseWidget
}

func New(child widget.BaseWidget, opts ...Option) widget.BaseWidget {
	link := &Link{
		BaseWidget: widget.NewWithOptions("a", map[string]any{
			"is": "pushstate-anchor",
		}),
	}

	opts = append(
		opts,
		Class("gf-link"),
	)

	for _, option := range opts {
		option()(link.BaseWidget)
	}

	link.AppendChild(child.Widget)
	return link.BaseWidget
}
