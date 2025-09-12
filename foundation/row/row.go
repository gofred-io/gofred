package row

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/theme"
)

type row struct {
	opts []div.Option
}

func New(children []application.BaseWidget, opts ...Option) application.BaseWidget {
	r := &row{}

	opts = append(
		opts,
		display(theme.DisplayTypeFlex),
		flexDirection(theme.FlexDirectionTypeRow),
	)

	for _, option := range opts {
		option(r)
	}

	return div.New(
		children,
		r.opts...,
	)
}
