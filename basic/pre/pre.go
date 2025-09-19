package pre

import (
	"github.com/gofred-io/gofred/application"
)

type pre struct {
	application.BaseWidget
}

func New(child application.BaseWidget, opts ...Option) application.BaseWidget {
	pre := &pre{
		BaseWidget: application.New("pre"),
	}

	for _, option := range opts {
		option()(pre.BaseWidget)
	}

	pre.AppendChild(child.Widget)
	return pre.BaseWidget
}
