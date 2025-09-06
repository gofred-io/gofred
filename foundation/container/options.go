package container

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/options/spacing"
)

func (container *FContainer) display(display DisplayType) *FContainer {
	container.BDiv.Display(display)
	return container
}

func (container *FContainer) BackgroundColor(color string) *FContainer {
	container.BDiv.BackgroundColor(color)
	return container
}

func (container *FContainer) BorderColor(color string) *FContainer {
	container.BDiv.BorderColor(color)
	return container
}

func (container *FContainer) BorderStyle(style BorderStyleType) *FContainer {
	container.BDiv.BorderStyle(style)
	return container
}

func (container *FContainer) BorderWidth(top int, right int, bottom int, left int) *FContainer {
	container.BDiv.BorderWidth(top, right, bottom, left)
	return container
}

func (container *FContainer) BorderRadius(borderRadius int) *FContainer {
	container.BDiv.BorderRadius(borderRadius)
	return container
}

func (container *FContainer) BoxShadow(shadow string) *FContainer {
	container.BDiv.BoxShadow(shadow)
	return container
}

func (container *FContainer) Class(class string) *FContainer {
	container.BDiv.Class(class)
	return container
}

func (container *FContainer) Flex(flex int) *FContainer {
	container.BDiv.Flex(flex)
	return container
}

func (container *FContainer) Height(opts ...BreakpointOptions[int]) *FContainer {
	container.BDiv.Height(opts...)
	return container
}

func (container *FContainer) ID(id string) *FContainer {
	container.BDiv.ID(id)
	return container
}

func (container *FContainer) Margin(opts ...BreakpointOptions[Spacing]) *FContainer {
	container.BDiv.Margin(opts...)
	return container
}

func (container *FContainer) MaxWidth(opts ...BreakpointOptions[int]) *FContainer {
	container.BDiv.MaxWidth(opts...)
	return container
}

func (container *FContainer) Overflow(overflow OverflowType) *FContainer {
	container.BDiv.Overflow(overflow)
	return container
}

func (container *FContainer) Padding(opts ...BreakpointOptions[Spacing]) *FContainer {
	container.BDiv.Padding(opts...)
	return container
}

func (container *FContainer) Visible(opts ...BreakpointOptions[bool]) *FContainer {
	container.BDiv.Visible(opts...)
	return container
}

func (container *FContainer) Width(opts ...BreakpointOptions[int]) *FContainer {
	container.BDiv.Width(opts...)
	return container
}

func (container *FContainer) WidthP(opts ...BreakpointOptions[float64]) *FContainer {
	container.BDiv.WidthP(opts...)
	return container
}
