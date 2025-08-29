package spacer

import (
	"github.com/gofred-io/gofred/container"
	"github.com/gofred-io/gofred/widget"
)

type Spacer struct {
	widget.Widget
	height int
	width  int
}

func New(options ...Options) widget.Widget {
	var (
		spacer       = &Spacer{}
		styleOptions = []container.StyleOptions{}
	)

	for _, option := range options {
		option(spacer)
	}

	if spacer.width > 0 {
		styleOptions = append(styleOptions, container.Width(spacer.width))
	}

	if spacer.height > 0 {
		styleOptions = append(styleOptions, container.Height(spacer.height))
	}

	if spacer.width == 0 && spacer.height == 0 {
		styleOptions = append(styleOptions, container.Flex(1))
	}

	return container.New(
		widget.Nil,
		container.Style(styleOptions...),
	)
}
