package spacer

import (
	"github.com/gofred-io/gofred/div"
	"github.com/gofred-io/gofred/widget"
)

type spacer struct {
	opts []div.Option
}

func New(opts ...Option) widget.BaseWidget {
	s := &spacer{}

	if len(opts) == 0 {
		opts = append(opts, flex(1))
	}

	for _, option := range opts {
		option(s)
	}

	return div.New(
		[]widget.BaseWidget{},
		s.opts...,
	)
}
