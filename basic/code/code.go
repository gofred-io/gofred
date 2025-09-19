package code

import (
	"github.com/gofred-io/gofred/application"
)

type code struct {
	application.BaseWidget
}

func New(innerCode string, opts ...Option) application.BaseWidget {
	code := &code{
		BaseWidget: application.New("code"),
	}

	for _, option := range opts {
		option()(code.BaseWidget)
	}

	code.Set("innerText", innerCode)
	return code.BaseWidget
}
