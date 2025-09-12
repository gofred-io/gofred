package container

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/theme"
)

type container struct {
	opts []div.Option
}

func New(child application.BaseWidget, opts ...Option) application.BaseWidget {
	var (
		c = &container{}
	)

	opts = append(
		opts,
		display(theme.DisplayTypeFlex),
	)

	for _, option := range opts {
		option(c)
	}

	return div.New(
		[]application.BaseWidget{child},
		c.opts...,
	)
}
