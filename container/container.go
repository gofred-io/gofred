package container

import (
	"github.com/gofred-io/gofred/div"
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/widget"
)

type Container struct {
	div.Div
	style   style.Style
	onClick func(widget widget.Widget)
}

func New(child widget.Widget, options ...Options) widget.Widget {
	container := &Container{
		style: style.Style{
			Display: &style.Display{
				Display: style.DisplayTypeFlex,
			},
		},
	}

	for _, option := range options {
		option(container)
	}

	container.Widget = div.New(
		[]widget.Widget{child},
		div.StyleFrom(&container.style),
	)

	container.SetOnClick(func(widget widget.Widget) {
		if container.onClick != nil {
			container.onClick(widget)
		}
	})

	return container.Widget
}
