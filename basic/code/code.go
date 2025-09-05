package code

import (
	"github.com/gofred-io/gofred/widget"
)

type code struct {
	*widget.BaseWidget
}

func New(innerCode string, opts ...Option) *code {
	code := &code{
		BaseWidget: widget.New("code"),
	}

	for _, option := range opts {
		option()(code)
	}

	code.Set("innerText", innerCode)
	return code
}
