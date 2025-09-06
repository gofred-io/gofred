package anchor

import (
	. "github.com/gofred-io/gofred/widget"
)

type BAnchor struct {
	*BaseWidget
}

func Anchor(child Widget) *BAnchor {
	anchor := &BAnchor{
		BaseWidget: NewWithOptions("a", map[string]any{
			"is": "pushstate-anchor",
		}),
	}

	anchor.AppendChild(child)
	return anchor
}
