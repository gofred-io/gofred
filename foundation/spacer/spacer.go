package spacer

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/widget"
)

type spacer struct {
	*widget.BaseWidget
	opts []div.Option
}

func New(opts ...Option) *spacer {
	s := &spacer{
		BaseWidget: &widget.BaseWidget{},
	}

	if len(opts) == 0 {
		opts = append(opts, flex(1))
	}

	for _, option := range opts {
		option(s)
	}

	s.Extend(
		div.New(
			[]widget.Widget{},
			s.opts...,
		),
	)
	return s
}
