package spacer

import (
	"github.com/gofred-io/gofred/container"
	"github.com/gofred-io/gofred/widget"
)

func New() widget.Widget {
	return container.New(
		widget.Nil,
		container.Style(
			container.Flex(1),
		),
	)
}
