package image

import (
	"github.com/gofred-io/gofred/widget"
)

type image struct {
	*widget.BaseWidget
}

func New(src string, opts ...Option) *image {
	image := &image{
		BaseWidget: widget.New("img"),
	}

	for _, option := range opts {
		option()(image)
	}

	image.SetAttribute("src", src)
	return image
}
