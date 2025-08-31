package row

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type row struct {
	opts []div.Option
}

func New(children []widget.BaseWidget, opts ...Option) widget.BaseWidget {
	r := &row{}

	opts = append(
		opts,
		display(options.DisplayTypeFlex),
		flexDirection(options.FlexDirectionTypeRow),
	)

	for _, option := range opts {
		option(r)
	}

	return div.New(
		children,
		r.opts...,
	)
}
