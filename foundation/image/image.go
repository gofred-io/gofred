package image

import (
	basicimage "github.com/gofred-io/gofred/basic/image"
	"github.com/gofred-io/gofred/widget"
)

type image struct {
	*widget.BaseWidget
	opts []basicimage.Option
}

func New(src string, opts ...Option) *image {
	image := &image{
		BaseWidget: &widget.BaseWidget{},
	}

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

	image.Extend(
		basicimage.New(src, image.opts...),
	)
	return image
}
