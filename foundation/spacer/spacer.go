package spacer

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/basic/div"
)

type spacer struct {
	opts []div.Option
}

func New(opts ...Option) application.BaseWidget {
	s := &spacer{}

	if len(opts) == 0 {
		opts = append(opts, flex(1))
	}

	for _, option := range opts {
		option(s)
	}

	return div.New(
		[]application.BaseWidget{},
		s.opts...,
	)
}
