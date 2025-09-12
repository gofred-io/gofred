package image

import (
	"github.com/gofred-io/gofred/application"
)

type Image struct {
	application.BaseWidget
}

func New(src string, opts ...Option) application.BaseWidget {
	image := &Image{
		BaseWidget: application.New("img"),
	}

	for _, option := range opts {
		option()(image.BaseWidget)
	}

	image.SetAttribute("src", src)
	return image.BaseWidget
}
