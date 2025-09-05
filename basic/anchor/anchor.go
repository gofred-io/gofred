package anchor

import "github.com/gofred-io/gofred/widget"

type anchor struct {
	*widget.BaseWidget
}

func New(child widget.Widget, options ...Option) *anchor {
	anchor := &anchor{
		BaseWidget: widget.NewWithOptions("a", map[string]any{
			"is": "pushstate-anchor",
		}),
	}

	for _, option := range options {
		option()(anchor)
	}

	anchor.AppendChild(child)
	return anchor
}
