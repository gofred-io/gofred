package container

import (
	"github.com/gofred-io/gofred/div"
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/widget"
)

type Container struct {
	div.Div
	style style.Style
}

func New(child widget.Widget, options ...Options) widget.Widget {
	container := &Container{}

	for _, option := range options {
		option(container)
	}

	return div.New(
		[]widget.Widget{child},
		div.StyleFrom(&container.style),
	)
}
