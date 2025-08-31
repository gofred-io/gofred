package anchor

import "github.com/gofred-io/gofred/widget"

type anchor struct {
	widget.BaseWidget
}

func New(child widget.BaseWidget, options ...Option) widget.BaseWidget {
	anchor := &anchor{
		BaseWidget: widget.NewWithOptions("a", map[string]any{
			"is": "pushstate-anchor",
		}),
	}

	for _, option := range options {
		option()(anchor.BaseWidget)
	}

	anchor.AppendChild(child.Widget)
	return anchor.BaseWidget
}
