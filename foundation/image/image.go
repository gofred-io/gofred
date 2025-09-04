package image

import (
	basicimage "github.com/gofred-io/gofred/basic/image"
	"github.com/gofred-io/gofred/widget"
)

type image struct {
	opts []basicimage.Option
}

func New(src string, opts ...Option) widget.BaseWidget {
	image := &image{}

	defaultOpts := []Option{
		Class("gf-image"),
	}

	opts = append(
		defaultOpts,
		opts...,
	)

	for _, option := range opts {
		option(image)
	}

	return basicimage.New(src, image.opts...)
}
