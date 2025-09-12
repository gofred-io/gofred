package span

import (
	"github.com/gofred-io/gofred/application"
)

type span struct {
	application.BaseWidget
}

func New(child application.BaseWidget, options ...Option) application.BaseWidget {
	span := &span{
		BaseWidget: application.New("span"),
	}

	for _, option := range options {
		option()(span.BaseWidget)
	}

	span.AppendChild(child.Widget)
	return span.BaseWidget
}
