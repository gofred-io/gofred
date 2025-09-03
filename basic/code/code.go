package code

import (
	"github.com/gofred-io/gofred/widget"
)

type code struct {
	widget.BaseWidget
}

func New(innerCode string, opts ...Option) widget.BaseWidget {
	code := &code{
		BaseWidget: widget.New("code"),
	}

	for _, option := range opts {
		option()(code.BaseWidget)
	}

	code.Set("innerText", innerCode)
	return code.BaseWidget
}
