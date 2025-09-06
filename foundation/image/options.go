package image

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/options"
)

func (image *FImage) Class(class string) *FImage {
	image.BImage.Class(class)
	return image
}

func (image *FImage) Height(opts ...BreakpointOptions[int]) *FImage {
	image.BImage.Height(opts...)
	return image
}

func (image *FImage) ID(id string) *FImage {
	image.BImage.ID(id)
	return image
}

func (image *FImage) MaxWidth(opts ...BreakpointOptions[int]) *FImage {
	image.BImage.MaxWidth(opts...)
	return image
}

func (image *FImage) UserSelect(userSelect UserSelectType) *FImage {
	image.BImage.UserSelect(userSelect)
	return image
}

func (image *FImage) Width(opts ...BreakpointOptions[int]) *FImage {
	image.BImage.Width(opts...)
	return image
}

func (image *FImage) WidthP(opts ...BreakpointOptions[float64]) *FImage {
	image.BImage.WidthP(opts...)
	return image
}
