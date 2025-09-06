package image

import (
	basicimage "github.com/gofred-io/gofred/basic/image"
)

type FImage struct {
	*basicimage.BImage
}

func Image(src string) *FImage {
	image := &FImage{
		BImage: basicimage.Image(src),
	}

	image.Class("gf-image")
	return image
}
