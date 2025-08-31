package column

import (
	"github.com/gofred-io/gofred/div"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

func New(children []widget.BaseWidget, opts ...Options) widget.BaseWidget {
	var (
		divOptions = []options.Options{}
	)

	for _, option := range opts {
		divOptions = append(divOptions, option())
	}

	divOptions = append(
		divOptions,
		options.Display(options.DisplayTypeFlex),
		options.FlexDirection(options.FlexDirectionTypeColumn),
	)

	return div.New(
		children,
		divOptions...,
	)
}
