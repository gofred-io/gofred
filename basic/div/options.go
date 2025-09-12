package div

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

type Option options.OptionWrapper

func AlignItems(alignItems theme.AxisAlignmentType) Option {
	return func() options.Option {
		return options.AlignItems(alignItems)
	}
}

func AlignSelf(alignSelf theme.AxisSizeType) Option {
	return func() options.Option {
		return options.AlignSelf(alignSelf)
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

func BorderColor(color string) Option {
	return func() options.Option {
		return options.BorderColor(color)
	}
}

func BorderStyle(style theme.BorderStyleType) Option {
	return func() options.Option {
		return options.BorderStyle(style)
	}
}

func BorderWidth(top int, right int, bottom int, left int) Option {
	return func() options.Option {
		return options.BorderWidth(top, right, bottom, left)
	}
}

func BorderRadius(borderRadius int) Option {
	return func() options.Option {
		return options.BorderRadius(borderRadius)
	}
}

func BoxShadow(shadow string) Option {
	return func() options.Option {
		return options.BoxShadow(shadow)
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

func Display(display theme.DisplayType) Option {
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

func FlexDirection(flexDirection theme.FlexDirectionType) Option {
	return func() options.Option {
		return options.FlexDirection(flexDirection)
	}
}

func FlexWrap(flexWrap theme.FlexWrapType) Option {
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

func JustifyContent(justifyContent theme.AxisAlignmentType) Option {
	return func() options.Option {
		return options.JustifyContent(justifyContent)
	}
}

func LineHeight(lineHeight float64) Option {
	return func() options.Option {
		return options.LineHeight(lineHeight)
	}
}

func Margin(opts ...breakpoint.BreakpointOptions[spacing.Spacing]) Option {
	return func() options.Option {
		return options.Margin(opts...)
	}
}

func MaxWidth(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func() options.Option {
		return options.MaxWidth(opts...)
	}
}

func NewTab(newTab bool) Option {
	return func() options.Option {
		return options.NewTab(newTab)
	}
}

func OnClick(handler options.OnClickHandler) Option {
	return func() options.Option {
		return options.OnClick(handler)
	}
}

func Overflow(overflow theme.OverflowType) Option {
	return func() options.Option {
		return options.Overflow(overflow)
	}
}

func Padding(opts ...breakpoint.BreakpointOptions[spacing.Spacing]) Option {
	return func() options.Option {
		return options.Padding(opts...)
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

func Transition(transition float64) Option {
	return func() options.Option {
		return options.Transition(transition)
	}
}

func UserSelect(userSelect theme.UserSelectType) Option {
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
