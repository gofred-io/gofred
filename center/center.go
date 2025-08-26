package center

import (
	"github.com/gofred-io/gofred/div"
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/widget"
)

type Center struct {
	div.Div
}

func New(ctx *widget.WidgetContext, options ...div.Options) widget.Widget {
	options = append(options, div.Style(
		div.Display(style.Display{
			Display:        style.DisplayTypeFlex,
			Flex:           1,
			AlignItems:     "center",
			JustifyContent: "center",
		}),
	))

	return div.New(
		ctx,
		options...,
	)
}
