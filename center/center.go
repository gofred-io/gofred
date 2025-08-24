package center

import (
	"github.com/man.go/mango/div"
	"github.com/man.go/mango/style"
	"github.com/man.go/mango/widget"
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
