package text

import (
	"github.com/gofred-io/gofred/basic/span"
	"github.com/gofred-io/gofred/widget"
)

type text struct {
	opts []span.Option
}

func New(innerText string, opts ...Option) widget.BaseWidget {
	text := &text{}

	opts = append(opts, setText(innerText))

	for _, option := range opts {
		option(text)
	}

	return span.New(widget.Nil, text.opts...)
}
