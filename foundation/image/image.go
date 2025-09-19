package image

import (
	"github.com/gofred-io/gofred/application"
	basicimage "github.com/gofred-io/gofred/basic/image"
)

type image struct {
	opts []basicimage.Option
}

func New(src string, opts ...Option) application.BaseWidget {
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
