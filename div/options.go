package div

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/options"
)

type Option options.OptionWrapper

func AlignItems(alignItems options.AxisAlignmentType) Option {
	return func() options.Option {
		return options.AlignItems(alignItems)
	}
}

func Alt(alt string) Option {
	return func() options.Option {
		return options.Alt(alt)
	}
}

func BackgroundColor(color string) Option {
	return func() options.Option {
		return options.BackgroundColor(color)
	}
}

func BorderRadius(borderRadius int) Option {
	return func() options.Option {
		return options.BorderRadius(borderRadius)
	}
}

func Class(class string) Option {
	return func() options.Option {
		return options.Class(class)
	}
}

func ColumnCount(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.ColumnCount(opts...)
	}
}

func ColumnGap(columnGap int) Option {
	return func() options.Option {
		return options.ColumnGap(columnGap)
	}
}

func D(data string) Option {
	return func() options.Option {
		return options.D(data)
	}
}

func Display(display options.DisplayType) Option {
	return func() options.Option {
		return options.Display(display)
	}
}

func Fill(fill string) Option {
	return func() options.Option {
		return options.Fill(fill)
	}
}

func Flex(flex int) Option {
	return func() options.Option {
		return options.Flex(flex)
	}
}

func FlexDirection(flexDirection options.FlexDirectionType) Option {
	return func() options.Option {
		return options.FlexDirection(flexDirection)
	}
}

func FlexWrap(flexWrap options.FlexWrapType) Option {
	return func() options.Option {
		return options.FlexWrap(flexWrap)
	}
}

func FontColor(color string) Option {
	return func() options.Option {
		return options.FontColor(color)
	}
}

func FontFamily(fontFamily string) Option {
	return func() options.Option {
		return options.FontFamily(fontFamily)
	}
}

func FontSize(fontSize int) Option {
	return func() options.Option {
		return options.FontSize(fontSize)
	}
}

func FontWeight(fontWeight string) Option {
	return func() options.Option {
		return options.FontWeight(fontWeight)
	}
}

func Height(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.Height(opts...)
	}
}

func Href(href string) Option {
	return func() options.Option {
		return options.Href(href)
	}
}

func ID(id string) Option {
	return func() options.Option {
		return options.ID(id)
	}
}

func JustifyContent(justifyContent options.AxisAlignmentType) Option {
	return func() options.Option {
		return options.JustifyContent(justifyContent)
	}
}

func LineHeight(lineHeight float64) Option {
	return func() options.Option {
		return options.LineHeight(lineHeight)
	}
}

func Margin(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.Margin(opts...)
	}
}

func MarginB(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.MarginB(opts...)
	}
}

func MarginH(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.MarginH(opts...)
	}
}

func MarginL(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.MarginL(opts...)
	}
}

func MarginR(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.MarginR(opts...)
	}
}

func MarginT(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.MarginT(opts...)
	}
}

func MarginV(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.MarginV(opts...)
	}
}

func MaxWidth(maxWidth int) Option {
	return func() options.Option {
		return options.MaxWidth(maxWidth)
	}
}

func OnClick(handler options.OnClickHandler) Option {
	return func() options.Option {
		return options.OnClick(handler)
	}
}

func NewTab(newTab bool) Option {
	return func() options.Option {
		return options.NewTab(newTab)
	}
}

func Padding(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.Padding(opts...)
	}
}

func PaddingB(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.PaddingB(opts...)
	}
}

func PaddingH(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.PaddingH(opts...)
	}
}

func PaddingL(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.PaddingL(opts...)
	}
}

func PaddingR(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.PaddingR(opts...)
	}
}

func PaddingT(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.PaddingT(opts...)
	}
}

func PaddingV(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.PaddingV(opts...)
	}
}

func RowGap(rowGap int) Option {
	return func() options.Option {
		return options.RowGap(rowGap)
	}
}

func Tooltip(tooltip string) Option {
	return func() options.Option {
		return options.Tooltip(tooltip)
	}
}

func UserSelect(userSelect options.UserSelectType) Option {
	return func() options.Option {
		return options.UserSelect(userSelect)
	}
}

func Visible(opts ...breakpoint.BreakpointOptions[bool]) Option {
	return func() options.Option {
		return options.Visible(opts...)
	}
}

func Width(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.Width(opts...)
	}
}

func WidthP(opts ...breakpoint.BreakpointOptions[float64]) Option {
	return func() options.Option {
		return options.WidthP(opts...)
	}
}
