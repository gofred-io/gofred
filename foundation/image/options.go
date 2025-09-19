package image

import (
	basicimage "github.com/gofred-io/gofred/basic/image"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/theme"
)

type Option func(image *image)

func Class(class string) Option {
	return func(image *image) {
		image.opts = append(image.opts, basicimage.Class(class))
	}
}

func Height(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(image *image) {
		image.opts = append(image.opts, basicimage.Height(opts...))
	}
}

func ID(id string) Option {
	return func(image *image) {
		image.opts = append(image.opts, basicimage.ID(id))
	}
}

func MaxWidth(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(image *image) {
		image.opts = append(image.opts, basicimage.MaxWidth(opts...))
	}
}

func UserSelect(userSelect theme.UserSelectType) Option {
	return func(image *image) {
		image.opts = append(image.opts, basicimage.UserSelect(userSelect))
	}
}

func Width(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(image *image) {
		image.opts = append(image.opts, basicimage.Width(opts...))
	}
}

func WidthP(opts ...breakpoint.BreakpointOptions[float64]) Option {
	return func(image *image) {
		image.opts = append(image.opts, basicimage.WidthP(opts...))
	}
}
