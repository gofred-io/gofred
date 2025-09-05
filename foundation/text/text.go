package text

import (
	"github.com/gofred-io/gofred/basic/span"
	"github.com/gofred-io/gofred/widget"
)

type text struct {
	*widget.BaseWidget
	opts []span.Option
}

func New(innerText string, opts ...Option) *text {
	text := &text{
		BaseWidget: &widget.BaseWidget{},
	}

	opts = append(opts, setText(innerText))

	for _, option := range opts {
		option(text)
	}

	text.Extend(
		span.New(nil, text.opts...),
	)

	return text
}
