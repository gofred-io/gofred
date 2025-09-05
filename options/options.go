package options

import (
	"fmt"
	"strconv"

	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/widget"
)

type Option func(widget widget.Widget)
type OptionWrapper func() Option
type OnClickHandler func(this widget.Widget, e widget.Event)

func AlignItems(alignItems AxisAlignmentType) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("align-items", string(alignItems))
	}
}

func AlignSelf(alignSelf AxisSizeType) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("align-self", string(alignSelf))
	}
}

func Alt(alt string) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().SetAttribute("alt", alt)
	}
}

func BackgroundColor(color string) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("background-color", color)
	}
}

func BorderColor(color string) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("border-color", color)
	}
}

func BorderStyle(style BorderStyleType) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("border-style", string(style))
	}
}

func BorderWidth(top int, right int, bottom int, left int) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("border-width", fmt.Sprintf("%dpx %dpx %dpx %dpx", top, right, bottom, left))
	}
}

func BorderRadius(borderRadius int) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("border-radius", fmt.Sprintf("%dpx", borderRadius))
	}
}

func BoxShadow(shadow string) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("box-shadow", shadow)
	}
}

func Class(class string) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().AddClass(class)
	}
}

func FontColor(color string) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("color", color)
	}
}

func ColumnCount(columnCount ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.Widget) {
		for _, option := range columnCount {
			option(widget.GetBaseWidget().ColumnCount)
		}

		listenBreakpointOption(widget.GetBaseWidget().ColumnCount, func(columnCount int) {
			widget.GetBaseWidget().UpdateStyleProperty("grid-template-columns", fmt.Sprintf("repeat(%d, 1fr)", columnCount))
		})
	}
}

func ColumnGap(columnGap int) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("column-gap", fmt.Sprintf("%dpx", columnGap))
	}
}

func D(data string) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().SetAttribute("d", data)
	}
}

func Display(display DisplayType) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("display", string(display))
	}
}

func Fill(fill string) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("fill", fill)
	}
}

func Flex(flex int) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("flex", strconv.Itoa(flex))
	}
}

func FlexDirection(flexDirection FlexDirectionType) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("flex-direction", string(flexDirection))
	}
}

func FlexWrap(flexWrap FlexWrapType) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("flex-wrap", string(flexWrap))
	}
}

func FontFamily(fontFamily string) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("font-family", fontFamily)
	}
}

func FontSize(fontSize int) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("font-size", fmt.Sprintf("%dpx", fontSize))
	}
}

func FontWeight(fontWeight string) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("font-weight", fontWeight)
	}
}

func Height(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.Widget) {
		for _, option := range options {
			option(widget.GetBaseWidget().Height)

			listenBreakpointOption(widget.GetBaseWidget().Height, func(height int) {
				widget.GetBaseWidget().UpdateStyleProperty("height", fmt.Sprintf("%dpx", height))
			})
		}
	}
}

func Href(href string) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().SetAttribute("href", href)
	}
}

func ID(id string) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().SetID(id)
	}
}

func JustifyContent(justifyContent AxisAlignmentType) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("justify-content", string(justifyContent))
	}
}

func LineHeight(lineHeight float64) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("line-height", fmt.Sprintf("%f", lineHeight))
	}
}

func Margin(options ...breakpoint.BreakpointOptions[spacing.Spacing]) Option {
	return func(widget widget.Widget) {
		for _, option := range options {
			option(widget.GetBaseWidget().Margin)
		}

		listenBreakpointOption(widget.GetBaseWidget().Margin, func(margin spacing.Spacing) {
			widget.GetBaseWidget().UpdateStyleProperty("margin", fmt.Sprintf("%dpx %dpx %dpx %dpx", margin.Top, margin.Right, margin.Bottom, margin.Left))
		})
	}
}

func MaxWidth(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.Widget) {
		for _, option := range options {
			option(widget.GetBaseWidget().MaxWidth)
		}

		listenBreakpointOption(widget.GetBaseWidget().MaxWidth, func(maxWidth int) {
			widget.GetBaseWidget().UpdateStyleProperty("max-width", fmt.Sprintf("%dpx", maxWidth))
		})
	}
}

func NewTab(newTab bool) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().SetAttribute("target", "_blank")
		if newTab {
			widget.GetBaseWidget().SetAttribute("target", "_blank")
		}
	}
}

func OnClick(handler OnClickHandler) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().SetOnClick(handler)
	}
}

func Overflow(overflow OverflowType) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("overflow", string(overflow))
	}
}

func Padding(options ...breakpoint.BreakpointOptions[spacing.Spacing]) Option {
	return func(widget widget.Widget) {
		for _, option := range options {
			option(widget.GetBaseWidget().Padding)
		}

		listenBreakpointOption(widget.GetBaseWidget().Padding, func(padding spacing.Spacing) {
			widget.GetBaseWidget().UpdateStyleProperty("padding", fmt.Sprintf("%dpx %dpx %dpx %dpx", padding.Top, padding.Right, padding.Bottom, padding.Left))
		})
	}
}

func RowGap(rowGap int) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("row-gap", fmt.Sprintf("%dpx", rowGap))
	}
}

func SetText(text string) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().SetText(text)
	}
}

func TextAlign(textAlign TextAlignType) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("text-align", string(textAlign))
	}
}

func Tooltip(tooltip string) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().SetAttribute("title", tooltip)
	}
}

func Transition(seconds float64) Option {
	return func(widget widget.Widget) {
		if seconds < 0 {
			seconds = 0.3
		}

		widget.GetBaseWidget().UpdateStyleProperty("transition", fmt.Sprintf("%.1fs", seconds))
	}
}

func UserSelect(userSelect UserSelectType) Option {
	return func(widget widget.Widget) {
		widget.GetBaseWidget().UpdateStyleProperty("user-select", string(userSelect))
	}
}

func Visible(options ...breakpoint.BreakpointOptions[bool]) Option {
	return func(widget widget.Widget) {
		for _, option := range options {
			option(widget.GetBaseWidget().Visible)
		}

		listenBreakpointOption(widget.GetBaseWidget().Visible, func(visible bool) {
			if visible {
				widget.GetBaseWidget().RemoveClass("gf-hidden")
			} else {
				widget.GetBaseWidget().AddClass("gf-hidden")
			}
		})
	}
}

func Width(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.Widget) {
		for _, option := range options {
			option(widget.GetBaseWidget().Width)
		}

		listenBreakpointOption(widget.GetBaseWidget().Width, func(width int) {
			widget.GetBaseWidget().UpdateStyleProperty("width", fmt.Sprintf("%dpx", width))
		})
	}
}

func WidthP(options ...breakpoint.BreakpointOptions[float64]) Option {
	return func(widget widget.Widget) {
		for _, option := range options {
			option(widget.GetBaseWidget().WidthP)
		}

		listenBreakpointOption(widget.GetBaseWidget().WidthP, func(width float64) {
			widget.GetBaseWidget().UpdateStyleProperty("width", fmt.Sprintf("%f%%", width*100))
		})
	}
}

func listenBreakpointOption[T any](breakpointValue *breakpoint.BreakpointValue[T], callback func(value T)) {
	breakpointHook := hooks.UseBreakpoint()
	innerCallback := func(breakPoint breakpoint.BreakPoint) {
		if breakpointValue == nil {
			return
		}

		value := breakpointValue.Get(breakPoint)
		callback(value)
	}

	listener := listenable.NewListener(innerCallback)
	breakpointHook.AddListener(listener)

	currentBreakpoint := breakpoint.GetCurrent()
	innerCallback(currentBreakpoint)
}
