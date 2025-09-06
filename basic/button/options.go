package button

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/options"
)

func (b *BButton) BackgroundColor(backgroundColor string) *BButton {
	BackgroundColor(backgroundColor)(b)
	return b
}

func (b *BButton) BorderRadius(borderRadius int) *BButton {
	BorderRadius(borderRadius)(b)
	return b
}

func (b *BButton) Class(class string) *BButton {
	Class(class)(b)
	return b
}

func (b *BButton) Height(opts ...BreakpointOptions[int]) *BButton {
	Height(opts...)(b)
	return b
}

func (b *BButton) ID(id string) *BButton {
	ID(id)(b)
	return b
}

func (b *BButton) MaxWidth(opts ...BreakpointOptions[int]) *BButton {
	MaxWidth(opts...)(b)
	return b
}

func (b *BButton) OnClick(handler OnClickHandler) *BButton {
	OnClick(handler)(b)
	return b
}

func (b *BButton) Tooltip(tooltip string) *BButton {
	Tooltip(tooltip)(b)
	return b
}

func (b *BButton) Visible(opts ...BreakpointOptions[bool]) *BButton {
	Visible(opts...)(b)
	return b
}

func (b *BButton) Width(opts ...BreakpointOptions[int]) *BButton {
	Width(opts...)(b)
	return b
}

func (b *BButton) WidthP(opts ...BreakpointOptions[float64]) *BButton {
	WidthP(opts...)(b)
	return b
}
