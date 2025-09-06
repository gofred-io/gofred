package button

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/options"
)

func (b *FButton) BackgroundColor(backgroundColor string) *FButton {
	b.BButton.BackgroundColor(backgroundColor)
	return b
}

func (b *FButton) BorderRadius(borderRadius int) *FButton {
	b.BButton.BorderRadius(borderRadius)
	return b
}

func (b *FButton) Class(class string) *FButton {
	b.BButton.Class(class)
	return b
}

func (b *FButton) Height(opts ...BreakpointOptions[int]) *FButton {
	b.BButton.Height(opts...)
	return b
}

func (b *FButton) ID(id string) *FButton {
	b.BButton.ID(id)
	return b
}

func (b *FButton) MaxWidth(opts ...BreakpointOptions[int]) *FButton {
	b.BButton.MaxWidth(opts...)
	return b
}

func (b *FButton) OnClick(handler OnClickHandler) *FButton {
	b.BButton.OnClick(handler)
	return b
}

func (b *FButton) Tooltip(tooltip string) *FButton {
	b.BButton.Tooltip(tooltip)
	return b
}

func (b *FButton) Visible(opts ...BreakpointOptions[bool]) *FButton {
	b.BButton.Visible(opts...)
	return b
}

func (b *FButton) Width(opts ...BreakpointOptions[int]) *FButton {
	b.BButton.Width(opts...)
	return b
}

func (b *FButton) WidthP(opts ...BreakpointOptions[float64]) *FButton {
	b.BButton.WidthP(opts...)
	return b
}
