package image

import (
	"github.com/gofred-io/gofred/widget"
)

type Image struct {
	widget.BaseWidget
}

func New(src string, opts ...Option) widget.BaseWidget {
	image := &Image{
		BaseWidget: widget.New("img"),
	}

	for _, option := range opts {
		option()(image.BaseWidget)
	}

	image.SetAttribute("src", src)
	return image.BaseWidget
}
