package icon

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/options"
)

func (icon *FIcon) Class(class string) *FIcon {
	icon.BSvg.Class(class)
	return icon
}

func (icon *FIcon) Fill(fill string) *FIcon {
	icon.BSvg.Fill(fill)
	return icon
}

func (icon *FIcon) Height(opts ...BreakpointOptions[int]) *FIcon {
	icon.BSvg.Height(opts...)
	return icon
}

func (icon *FIcon) ID(id string) *FIcon {
	icon.BSvg.ID(id)
	return icon
}

func (icon *FIcon) MaxWidth(opts ...BreakpointOptions[int]) *FIcon {
	icon.BSvg.MaxWidth(opts...)
	return icon
}

func (icon *FIcon) UserSelect(userSelect UserSelectType) *FIcon {
	icon.BSvg.UserSelect(userSelect)
	return icon
}

func (icon *FIcon) Width(opts ...BreakpointOptions[int]) *FIcon {
	icon.BSvg.Width(opts...)
	return icon
}

func (icon *FIcon) WidthP(opts ...BreakpointOptions[float64]) *FIcon {
	icon.BSvg.WidthP(opts...)
	return icon
}
