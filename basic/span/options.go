package span

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/options/spacing"
)

func (s *BSpan) AlignItems(alignItems AxisAlignmentType) *BSpan {
	AlignItems(alignItems)(s)
	return s
}

func (s *BSpan) Alt(alt string) *BSpan {
	Alt(alt)(s)
	return s
}

func (s *BSpan) BackgroundColor(color string) *BSpan {
	BackgroundColor(color)(s)
	return s
}

func (s *BSpan) BorderRadius(borderRadius int) *BSpan {
	BorderRadius(borderRadius)(s)
	return s
}

func (s *BSpan) Class(class string) *BSpan {
	Class(class)(s)
	return s
}

func (s *BSpan) ColumnCount(opts ...BreakpointOptions[int]) *BSpan {
	ColumnCount(opts...)(s)
	return s
}

func (s *BSpan) ColumnGap(columnGap int) *BSpan {
	ColumnGap(columnGap)(s)
	return s
}

func (s *BSpan) D(data string) *BSpan {
	D(data)(s)
	return s
}

func (s *BSpan) Display(display DisplayType) *BSpan {
	Display(display)(s)
	return s
}

func (s *BSpan) Fill(fill string) *BSpan {
	Fill(fill)(s)
	return s
}

func (s *BSpan) Flex(flex int) *BSpan {
	Flex(flex)(s)
	return s
}

func (s *BSpan) FlexDirection(flexDirection FlexDirectionType) *BSpan {
	FlexDirection(flexDirection)(s)
	return s
}

func (s *BSpan) FlexWrap(flexWrap FlexWrapType) *BSpan {
	FlexWrap(flexWrap)(s)
	return s
}

func (s *BSpan) FontColor(color string) *BSpan {
	FontColor(color)(s)
	return s
}

func (s *BSpan) FontFamily(fontFamily string) *BSpan {
	FontFamily(fontFamily)(s)
	return s
}

func (s *BSpan) FontSize(fontSize int) *BSpan {
	FontSize(fontSize)(s)
	return s
}

func (s *BSpan) FontWeight(fontWeight string) *BSpan {
	FontWeight(fontWeight)(s)
	return s
}

func (s *BSpan) Height(opts ...BreakpointOptions[int]) *BSpan {
	Height(opts...)(s)
	return s
}

func (s *BSpan) Href(href string) *BSpan {
	Href(href)(s)
	return s
}

func (s *BSpan) ID(id string) *BSpan {
	ID(id)(s)
	return s
}

func (s *BSpan) JustifyContent(justifyContent AxisAlignmentType) *BSpan {
	JustifyContent(justifyContent)(s)
	return s
}

func (s *BSpan) LineHeight(lineHeight float64) *BSpan {
	LineHeight(lineHeight)(s)
	return s
}

func (s *BSpan) Margin(opts ...BreakpointOptions[Spacing]) *BSpan {
	Margin(opts...)(s)
	return s
}

func (s *BSpan) MaxWidth(opts ...BreakpointOptions[int]) *BSpan {
	MaxWidth(opts...)(s)
	return s
}

func (s *BSpan) OnClick(handler OnClickHandler) *BSpan {
	OnClick(handler)(s)
	return s
}

func (s *BSpan) NewTab(newTab bool) *BSpan {
	NewTab(newTab)(s)
	return s
}

func (s *BSpan) Padding(opts ...BreakpointOptions[Spacing]) *BSpan {
	Padding(opts...)(s)
	return s
}

func (s *BSpan) RowGap(rowGap int) *BSpan {
	RowGap(rowGap)(s)
	return s
}

func (s *BSpan) SetText(text string) *BSpan {
	SetText(text)(s)
	return s
}

func (s *BSpan) TextAlign(textAlign TextAlignType) *BSpan {
	TextAlign(textAlign)(s)
	return s
}

func (s *BSpan) Tooltip(tooltip string) *BSpan {
	Tooltip(tooltip)(s)
	return s
}

func (s *BSpan) UserSelect(userSelect UserSelectType) *BSpan {
	UserSelect(userSelect)(s)
	return s
}

func (s *BSpan) Visible(opts ...BreakpointOptions[bool]) *BSpan {
	Visible(opts...)(s)
	return s
}

func (s *BSpan) Width(opts ...BreakpointOptions[int]) *BSpan {
	Width(opts...)(s)
	return s
}

func (s *BSpan) WidthP(opts ...BreakpointOptions[float64]) *BSpan {
	WidthP(opts...)(s)
	return s
}
