package image

import (
	. "github.com/gofred-io/gofred/widget"
)

type BImage struct {
	*BaseWidget
}

func Image(src string) *BImage {
	image := &BImage{
		BaseWidget: New("img"),
	}

	image.SetAttribute("src", src)
	return image
}
