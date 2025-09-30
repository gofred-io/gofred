package options

import (
	"fmt"
	"strconv"

	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

type Option func(widget application.BaseWidget)
type OptionWrapper func() Option
type OnClickHandler func(this application.BaseWidget, e application.Event)
type OnDragEndHandler func(this application.BaseWidget, e application.Event)
type OnDragHandler func(this application.BaseWidget, e application.Event)
type OnDragStartHandler func(this application.BaseWidget, e application.Event)

func AlignItems(alignItems theme.AxisAlignmentType) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("align-items", string(alignItems))
	}
}

func AlignSelf(alignSelf theme.AxisSizeType) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("align-self", string(alignSelf))
	}
}

func Alt(alt string) Option {
	return func(widget application.BaseWidget) {
		widget.SetAttribute("alt", alt)
	}
}

func BackgroundColor(color string) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("background-color", color)
	}
}

func BorderColor(color string) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("border-color", color)
	}
}

func BorderStyle(style theme.BorderStyleType) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("border-style", string(style))
	}
}

func BorderWidth(spacing spacing.Spacing) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("border-width", fmt.Sprintf("%dpx %dpx %dpx %dpx", spacing.Top, spacing.Right, spacing.Bottom, spacing.Left))
	}
}

func BorderRadius(borderRadius int) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("border-radius", fmt.Sprintf("%dpx", borderRadius))
	}
}

func BoxShadow(shadow string) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("box-shadow", shadow)
	}
}

func Class(class string) Option {
	return func(widget application.BaseWidget) {
		widget.AddClass(class)
	}
}

func FontColor(color string) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("color", color)
	}
}

func ColumnCount(columnCount ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget application.BaseWidget) {
		for _, option := range columnCount {
			option(widget.ColumnCount)
		}

		listenBreakpointOption(widget.ColumnCount, func(columnCount int) {
			widget.UpdateStyleProperty("grid-template-columns", fmt.Sprintf("repeat(%d, 1fr)", columnCount))
		})
	}
}

func ColumnGap(columnGap int) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("column-gap", fmt.Sprintf("%dpx", columnGap))
	}
}

func D(data string) Option {
	return func(widget application.BaseWidget) {
		widget.SetAttribute("d", data)
	}
}

func Display(display theme.DisplayType) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("display", string(display))
	}
}

func Draggable(draggable bool) Option {
	return func(widget application.BaseWidget) {
		widget.SetDraggable(draggable)
	}
}

func Fill(fill string) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("fill", fill)
	}
}

func Flex(flex int) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("flex", strconv.Itoa(flex))
	}
}

func FlexDirection(flexDirection theme.FlexDirectionType) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("flex-direction", string(flexDirection))
	}
}

func FlexWrap(flexWrap theme.FlexWrapType) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("flex-wrap", string(flexWrap))
	}
}

func FontFamily(fontFamily string) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("font-family", fontFamily)
	}
}

func FontSize(fontSize int) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("font-size", fmt.Sprintf("%dpx", fontSize))
	}
}

func FontWeight(fontWeight string) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("font-weight", fontWeight)
	}
}

func Height(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget application.BaseWidget) {
		for _, option := range options {
			option(widget.Height)

			listenBreakpointOption(widget.Height, func(height int) {
				widget.UpdateStyleProperty("height", fmt.Sprintf("%dpx", height))
			})
		}
	}
}

func Href(href string) Option {
	return func(widget application.BaseWidget) {
		widget.SetAttribute("href", href)
	}
}

func ID(id string) Option {
	return func(widget application.BaseWidget) {
		widget.SetID(id)
	}
}

func Label(label string) Option {
	return func(widget application.BaseWidget) {
		widget.SetAttribute("aria-label", label)
	}
}

func JustifyContent(justifyContent theme.AxisAlignmentType) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("justify-content", string(justifyContent))
	}
}

func LetterSpacing(letterSpacing float64) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("letter-spacing", fmt.Sprintf("%f", letterSpacing))
	}
}

func LineHeight(lineHeight float64) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("line-height", fmt.Sprintf("%f", lineHeight))
	}
}

func Margin(options ...breakpoint.BreakpointOptions[spacing.Spacing]) Option {
	return func(widget application.BaseWidget) {
		for _, option := range options {
			option(widget.Margin)
		}

		listenBreakpointOption(widget.Margin, func(margin spacing.Spacing) {
			widget.UpdateStyleProperty("margin", fmt.Sprintf("%dpx %dpx %dpx %dpx", margin.Top, margin.Right, margin.Bottom, margin.Left))
		})
	}
}

func MaxWidth(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget application.BaseWidget) {
		for _, option := range options {
			option(widget.MaxWidth)
		}

		listenBreakpointOption(widget.MaxWidth, func(maxWidth int) {
			widget.UpdateStyleProperty("max-width", fmt.Sprintf("%dpx", maxWidth))
		})
	}
}

func NewTab(newTab bool) Option {
	return func(widget application.BaseWidget) {
		if newTab {
			widget.SetAttribute("target", "_blank")
		}
	}
}

func OnClick(handler OnClickHandler) Option {
	return func(widget application.BaseWidget) {
		widget.SetOnClick(handler)
	}
}

func OnDragEnd(handler OnDragEndHandler) Option {
	return func(widget application.BaseWidget) {
		widget.SetOnDragEnd(handler)
	}
}

func OnDrag(handler OnDragHandler) Option {
	return func(widget application.BaseWidget) {
		widget.SetOnDrag(handler)
	}
}

func OnDragStart(handler OnDragStartHandler) Option {
	return func(widget application.BaseWidget) {
		widget.SetOnDragStart(handler)
	}
}

func Opacity(opacity float64) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("opacity", fmt.Sprintf("%f", opacity))
	}
}

func Overflow(overflow theme.OverflowType) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("overflow", string(overflow))
	}
}

func Padding(options ...breakpoint.BreakpointOptions[spacing.Spacing]) Option {
	return func(widget application.BaseWidget) {
		for _, option := range options {
			option(widget.Padding)
		}

		listenBreakpointOption(widget.Padding, func(padding spacing.Spacing) {
			widget.UpdateStyleProperty("padding", fmt.Sprintf("%dpx %dpx %dpx %dpx", padding.Top, padding.Right, padding.Bottom, padding.Left))
		})
	}
}

func RowGap(rowGap int) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("row-gap", fmt.Sprintf("%dpx", rowGap))
	}
}

func SetText(text string) Option {
	return func(widget application.BaseWidget) {
		widget.SetText(text)
	}
}

func TextAlign(textAlign theme.TextAlignType) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("text-align", string(textAlign))
	}
}

func Tooltip(tooltip string) Option {
	return func(widget application.BaseWidget) {
		widget.SetAttribute("title", tooltip)
	}
}

func Transition(seconds float64) Option {
	return func(widget application.BaseWidget) {
		if seconds < 0 {
			seconds = 0.3
		}

		widget.UpdateStyleProperty("transition", fmt.Sprintf("%.1fs", seconds))
	}
}

func UserSelect(userSelect theme.UserSelectType) Option {
	return func(widget application.BaseWidget) {
		widget.UpdateStyleProperty("user-select", string(userSelect))
	}
}

func Visible(options ...breakpoint.BreakpointOptions[bool]) Option {
	return func(widget application.BaseWidget) {
		for _, option := range options {
			option(widget.Visible)
		}

		listenBreakpointOption(widget.Visible, func(visible bool) {
			if visible {
				widget.RemoveClass("gf-hidden")
			} else {
				widget.AddClass("gf-hidden")
			}
		})
	}
}

func Width(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget application.BaseWidget) {
		for _, option := range options {
			option(widget.Width)
		}

		listenBreakpointOption(widget.Width, func(width int) {
			widget.UpdateStyleProperty("width", fmt.Sprintf("%dpx", width))
		})
	}
}

func WidthP(options ...breakpoint.BreakpointOptions[float64]) Option {
	return func(widget application.BaseWidget) {
		for _, option := range options {
			option(widget.WidthP)
		}

		listenBreakpointOption(widget.WidthP, func(width float64) {
			widget.UpdateStyleProperty("width", fmt.Sprintf("%f%%", width*100))
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
