package button

import (
	"github.com/gofred-io/gofred/application"
)

type button struct {
	application.BaseWidget
}

func New(child application.BaseWidget, options ...Option) application.BaseWidget {
	button := &button{
		BaseWidget: application.New("button"),
	}

	for _, opt := range options {
		opt()(button.BaseWidget)
	}

	button.AppendChild(child.Widget)

	return button.BaseWidget
}
