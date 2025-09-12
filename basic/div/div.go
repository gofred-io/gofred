package div

import (
	"github.com/gofred-io/gofred/application"
)

type div struct {
	application.BaseWidget
}

func New(children []application.BaseWidget, opts ...Option) application.BaseWidget {
	div := &div{
		BaseWidget: application.New("div"),
	}

	for _, option := range opts {
		option()(div.BaseWidget)
	}

	for _, child := range children {
		div.AppendChild(child.Widget)
	}

	return div.BaseWidget
}
