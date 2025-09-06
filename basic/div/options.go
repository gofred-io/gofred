package div

import (
	"github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/options/spacing"
)

func (d *BDiv) AlignItems(alignItems AxisAlignmentType) *BDiv {
	AlignItems(alignItems)(d)
	return d
}

func (d *BDiv) AlignSelf(alignSelf AxisSizeType) *BDiv {
	AlignSelf(alignSelf)(d)
	return d
}

func (d *BDiv) Alt(alt string) *BDiv {
	Alt(alt)(d)
	return d
}

func (d *BDiv) BackgroundColor(color string) *BDiv {
	BackgroundColor(color)(d)
	return d
}

func (d *BDiv) BorderColor(color string) *BDiv {
	BorderColor(color)(d)
	return d
}

func (d *BDiv) BorderStyle(style BorderStyleType) *BDiv {
	BorderStyle(style)(d)
	return d
}

func (d *BDiv) BorderWidth(top int, right int, bottom int, left int) *BDiv {
	BorderWidth(top, right, bottom, left)(d)
	return d
}

func (d *BDiv) BorderRadius(borderRadius int) *BDiv {
	BorderRadius(borderRadius)(d)
	return d
}

func (d *BDiv) BoxShadow(shadow string) *BDiv {
	BoxShadow(shadow)(d)
	return d
}

func (d *BDiv) Class(class string) *BDiv {
	Class(class)(d)
	return d
}

func (d *BDiv) ColumnCount(opts ...breakpoint.BreakpointOptions[int]) *BDiv {
	ColumnCount(opts...)(d)
	return d
}

func (d *BDiv) ColumnGap(columnGap int) *BDiv {
	ColumnGap(columnGap)(d)
	return d
}

func (d *BDiv) D(data string) *BDiv {
	D(data)(d)
	return d
}

func (d *BDiv) Display(display DisplayType) *BDiv {
	Display(display)(d)
	return d
}

func (d *BDiv) Fill(fill string) *BDiv {
	Fill(fill)(d)
	return d
}

func (d *BDiv) Flex(flex int) *BDiv {
	Flex(flex)(d)
	return d
}

func (d *BDiv) FlexDirection(flexDirection FlexDirectionType) *BDiv {
	FlexDirection(flexDirection)(d)
	return d
}

func (d *BDiv) FlexWrap(flexWrap FlexWrapType) *BDiv {
	FlexWrap(flexWrap)(d)
	return d
}

func (d *BDiv) FontColor(color string) *BDiv {
	FontColor(color)(d)
	return d
}

func (d *BDiv) FontFamily(fontFamily string) *BDiv {
	FontFamily(fontFamily)(d)
	return d
}

func (d *BDiv) FontSize(fontSize int) *BDiv {
	FontSize(fontSize)(d)
	return d
}

func (d *BDiv) FontWeight(fontWeight string) *BDiv {
	FontWeight(fontWeight)(d)
	return d
}

func (d *BDiv) Height(opts ...breakpoint.BreakpointOptions[int]) *BDiv {
	Height(opts...)(d)
	return d
}

func (d *BDiv) Href(href string) *BDiv {
	Href(href)(d)
	return d
}

func (d *BDiv) ID(id string) *BDiv {
	ID(id)(d)
	return d
}

func (d *BDiv) JustifyContent(justifyContent AxisAlignmentType) *BDiv {
	JustifyContent(justifyContent)(d)
	return d
}

func (d *BDiv) LineHeight(lineHeight float64) *BDiv {
	LineHeight(lineHeight)(d)
	return d
}

func (d *BDiv) Margin(opts ...breakpoint.BreakpointOptions[Spacing]) *BDiv {
	Margin(opts...)(d)
	return d
}

func (d *BDiv) MaxWidth(opts ...breakpoint.BreakpointOptions[int]) *BDiv {
	MaxWidth(opts...)(d)
	return d
}

func (d *BDiv) NewTab(newTab bool) *BDiv {
	NewTab(newTab)(d)
	return d
}

func (d *BDiv) NoFlex() *BDiv {
	NoFlex()(d)
	return d
}

func (d *BDiv) OnClick(handler OnClickHandler) *BDiv {
	OnClick(handler)(d)
	return d
}

func (d *BDiv) Overflow(overflow OverflowType) *BDiv {
	Overflow(overflow)(d)
	return d
}

func (d *BDiv) Padding(opts ...breakpoint.BreakpointOptions[Spacing]) *BDiv {
	Padding(opts...)(d)
	return d
}

func (d *BDiv) RowGap(rowGap int) *BDiv {
	RowGap(rowGap)(d)
	return d
}

func (d *BDiv) Tooltip(tooltip string) *BDiv {
	Tooltip(tooltip)(d)
	return d
}

func (d *BDiv) Transition(transition float64) *BDiv {
	Transition(transition)(d)
	return d
}

func (d *BDiv) UserSelect(userSelect UserSelectType) *BDiv {
	UserSelect(userSelect)(d)
	return d
}

func (d *BDiv) Visible(opts ...breakpoint.BreakpointOptions[bool]) *BDiv {
	Visible(opts...)(d)
	return d
}

func (d *BDiv) Width(opts ...breakpoint.BreakpointOptions[int]) *BDiv {
	Width(opts...)(d)
	return d
}

func (d *BDiv) WidthP(opts ...breakpoint.BreakpointOptions[float64]) *BDiv {
	WidthP(opts...)(d)
	return d
}
