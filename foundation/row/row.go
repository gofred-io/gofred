package row

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type row struct {
	*widget.BaseWidget
	opts []div.Option
}

func New(children []widget.Widget, opts ...Option) *row {
	r := &row{
		BaseWidget: &widget.BaseWidget{},
	}

	opts = append(
		opts,
		display(options.DisplayTypeFlex),
		flexDirection(options.FlexDirectionTypeRow),
	)

	for _, option := range opts {
		option(r)
	}

	r.Extend(
		div.New(
			children,
			r.opts...,
		),
	)

	return r
}
