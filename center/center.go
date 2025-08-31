package center

import (
	"github.com/gofred-io/gofred/div"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type Center struct {
	widget.BaseWidget
}

func New(child widget.BaseWidget, opts ...options.Options) widget.BaseWidget {
	opts = append(
		opts,
		options.Display(options.DisplayTypeFlex),
		options.Flex(1),
		alignItems(options.AxisAlignmentTypeCenter),
		justifyContent(options.AxisAlignmentTypeCenter),
	)

	return div.New(
		[]widget.BaseWidget{child},
		opts...,
	)
}
