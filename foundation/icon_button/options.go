package iconbutton

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/options"
)

func (iconButton *FIconButton) BorderRadius(borderRadius int) *FIconButton {
	iconButton.BButton.BorderRadius(borderRadius)
	return iconButton
}

func (iconButton *FIconButton) Class(class string) *FIconButton {
	iconButton.BButton.Class(class)
	return iconButton
}

func (iconButton *FIconButton) Fill(fill string) *FIconButton {
	iconButton.FIcon.Fill(fill)
	return iconButton
}

func (iconButton *FIconButton) Height(opts ...BreakpointOptions[int]) *FIconButton {
	iconButton.BButton.Height(opts...)
	return iconButton
}

func (iconButton *FIconButton) ID(id string) *FIconButton {
	iconButton.BButton.ID(id)
	return iconButton
}

func (iconButton *FIconButton) MaxWidth(opts ...BreakpointOptions[int]) *FIconButton {
	iconButton.BButton.MaxWidth(opts...)
	return iconButton
}

func (iconButton *FIconButton) OnClick(handler OnClickHandler) *FIconButton {
	iconButton.BButton.OnClick(handler)
	return iconButton
}

func (iconButton *FIconButton) Tooltip(tooltip string) *FIconButton {
	iconButton.BButton.Tooltip(tooltip)
	return iconButton
}

func (iconButton *FIconButton) UserSelect(userSelect UserSelectType) *FIconButton {
	iconButton.FIcon.UserSelect(userSelect)
	return iconButton
}
func (iconButton *FIconButton) Visible(opts ...BreakpointOptions[bool]) *FIconButton {
	iconButton.BButton.Visible(opts...)
	return iconButton
}

func (iconButton *FIconButton) Width(opts ...BreakpointOptions[int]) *FIconButton {
	iconButton.BButton.Width(opts...)
	return iconButton
}

func (iconButton *FIconButton) WidthP(opts ...BreakpointOptions[float64]) *FIconButton {
	iconButton.BButton.WidthP(opts...)
	return iconButton
}
