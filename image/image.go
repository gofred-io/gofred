package image

import "github.com/gofred-io/gofred/widget"

type Image struct {
	widget.Widget
}

func New(src string, options ...Options) widget.Widget {
	image := &Image{
		Widget: widget.Context().CreateElement("img"),
	}

	for _, option := range options {
		option(image)
	}

	image.Set("src", src)
	return image.Widget
}
