package column

import (
	"github.com/gofred-io/gofred/div"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

func New(children []widget.BaseWidget, opts ...options.Options) widget.BaseWidget {
	opts = append(
		opts,
		options.Display(options.DisplayTypeFlex),
		options.FlexDirection(options.FlexDirectionTypeColumn),
		options.Flex(1),
	)

	return div.New(
		children,
		opts...,
	)
}
