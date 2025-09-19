package anchor

import (
	"github.com/gofred-io/gofred/application"
)

type anchor struct {
	application.BaseWidget
}

func New(child application.BaseWidget, options ...Option) application.BaseWidget {
	anchor := &anchor{
		BaseWidget: application.NewWithOptions("a", map[string]any{
			"is": "pushstate-anchor",
		}),
	}

	for _, option := range options {
		option()(anchor.BaseWidget)
	}

	anchor.AppendChild(child.Widget)
	return anchor.BaseWidget
}
