package image

import (
	"github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/options"
)

func (i *BImage) Class(class string) *BImage {
	Class(class)(i)
	return i
}

func (i *BImage) Height(opts ...breakpoint.BreakpointOptions[int]) *BImage {
	Height(opts...)(i)
	return i
}

func (i *BImage) ID(id string) *BImage {
	ID(id)(i)
	return i
}

func (i *BImage) MaxWidth(opts ...breakpoint.BreakpointOptions[int]) *BImage {
	MaxWidth(opts...)(i)
	return i
}

func (i *BImage) UserSelect(userSelect UserSelectType) *BImage {
	UserSelect(userSelect)(i)
	return i
}

func (i *BImage) Width(opts ...breakpoint.BreakpointOptions[int]) *BImage {
	Width(opts...)(i)
	return i
}

func (i *BImage) WidthP(opts ...breakpoint.BreakpointOptions[float64]) *BImage {
	WidthP(opts...)(i)
	return i
}
