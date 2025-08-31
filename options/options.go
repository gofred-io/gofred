package options

import (
	"fmt"
	"strconv"

	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/widget"
)

type Option func(widget widget.BaseWidget)
type OptionWrapper func() Option
type OnClickHandler func(this widget.BaseWidget, e widget.Event)

func AlignItems(alignItems AxisAlignmentType) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("align-items", string(alignItems))
	}
}

func Alt(alt string) Option {
	return func(widget widget.BaseWidget) {
		widget.SetAttribute("alt", alt)
	}
}

func BackgroundColor(color string) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("background-color", color)
	}
}

func BorderRadius(borderRadius int) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("border-radius", fmt.Sprintf("%dpx", borderRadius))
	}
}

func Class(class string) Option {
	return func(widget widget.BaseWidget) {
		widget.AddClass(class)
	}
}

func FontColor(color string) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("color", color)
	}
}

func ColumnCount(columnCount ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.BaseWidget) {
		for _, option := range columnCount {
			option(widget.ColumnCount)
		}

		listenBreakpointOption(widget.ColumnCount, func(columnCount int) {
			widget.UpdateStyleProperty("grid-template-columns", fmt.Sprintf("repeat(%d, auto)", columnCount))
		})
	}
}

func ColumnGap(columnGap int) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("column-gap", fmt.Sprintf("%dpx", columnGap))
	}
}

func D(data string) Option {
	return func(widget widget.BaseWidget) {
		widget.SetAttribute("d", data)
	}
}

func Display(display DisplayType) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("display", string(display))
	}
}

func Fill(fill string) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("fill", fill)
	}
}

func Flex(flex int) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("flex", strconv.Itoa(flex))
	}
}

func FlexDirection(flexDirection FlexDirectionType) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("flex-direction", string(flexDirection))
	}
}

func FlexWrap(flexWrap FlexWrapType) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("flex-wrap", string(flexWrap))
	}
}

func FontFamily(fontFamily string) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("font-family", fontFamily)
	}
}

func FontSize(fontSize int) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("font-size", fmt.Sprintf("%dpx", fontSize))
	}
}

func FontWeight(fontWeight string) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("font-weight", fontWeight)
	}
}

func Height(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Height)

			listenBreakpointOption(widget.Height, func(height int) {
				widget.UpdateStyleProperty("height", fmt.Sprintf("%dpx", height))
			})
		}
	}
}

func Href(href string) Option {
	return func(widget widget.BaseWidget) {
		widget.SetAttribute("href", href)
	}
}

func ID(id string) Option {
	return func(widget widget.BaseWidget) {
		widget.SetID(id)
	}
}

func JustifyContent(justifyContent AxisAlignmentType) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("justify-content", string(justifyContent))
	}
}

func LineHeight(lineHeight float64) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("line-height", fmt.Sprintf("%f", lineHeight))
	}
}

func Margin(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Margin)
		}

		listenBreakpointOption(widget.Margin, func(margin int) {
			widget.UpdateStyleProperty("margin", fmt.Sprintf("%dpx", margin))
		})
	}
}

func MarginB(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Margin)
		}

		listenBreakpointOption(widget.Margin, func(margin int) {
			widget.UpdateStyleProperty("margin-bottom", fmt.Sprintf("%dpx", margin))
		})
	}
}

func MarginH(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Margin)
		}

		listenBreakpointOption(widget.Margin, func(margin int) {
			widget.UpdateStyleProperty("margin-left", fmt.Sprintf("%dpx", margin))
			widget.UpdateStyleProperty("margin-right", fmt.Sprintf("%dpx", margin))
		})
	}
}

func MarginL(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Margin)
		}

		listenBreakpointOption(widget.Margin, func(margin int) {
			widget.UpdateStyleProperty("margin-left", fmt.Sprintf("%dpx", margin))
		})
	}
}

func MarginR(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Margin)
		}

		listenBreakpointOption(widget.Margin, func(margin int) {
			widget.UpdateStyleProperty("margin-right", fmt.Sprintf("%dpx", margin))
		})
	}
}

func MarginT(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Margin)
		}

		listenBreakpointOption(widget.Margin, func(margin int) {
			widget.UpdateStyleProperty("margin-top", fmt.Sprintf("%dpx", margin))
		})
	}
}

func MarginV(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Margin)
		}

		listenBreakpointOption(widget.Margin, func(margin int) {
			widget.UpdateStyleProperty("margin-top", fmt.Sprintf("%dpx", margin))
			widget.UpdateStyleProperty("margin-bottom", fmt.Sprintf("%dpx", margin))
		})
	}
}

func MaxWidth(maxWidth int) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("max-width", fmt.Sprintf("%dpx", maxWidth))
	}
}

func OnClick(handler OnClickHandler) Option {
	return func(widget widget.BaseWidget) {
		widget.SetOnClick(handler)
	}
}

func NewTab(newTab bool) Option {
	return func(widget widget.BaseWidget) {
		if newTab {
			widget.SetAttribute("target", "_blank")
		}
	}
}

func Padding(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Padding)
		}

		listenBreakpointOption(widget.Padding, func(padding int) {
			widget.UpdateStyleProperty("padding", fmt.Sprintf("%dpx", padding))
		})
	}
}

func PaddingB(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Padding)
		}

		listenBreakpointOption(widget.Padding, func(padding int) {
			widget.UpdateStyleProperty("padding-bottom", fmt.Sprintf("%dpx", padding))
		})
	}
}

func PaddingH(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Padding)
		}

		listenBreakpointOption(widget.Padding, func(padding int) {
			widget.UpdateStyleProperty("padding-left", fmt.Sprintf("%dpx", padding))
			widget.UpdateStyleProperty("padding-right", fmt.Sprintf("%dpx", padding))
		})
	}
}

func PaddingL(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Padding)
		}

		listenBreakpointOption(widget.Padding, func(padding int) {
			widget.UpdateStyleProperty("padding-left", fmt.Sprintf("%dpx", padding))
		})
	}
}

func PaddingR(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Padding)
		}

		listenBreakpointOption(widget.Padding, func(padding int) {
			widget.UpdateStyleProperty("padding-right", fmt.Sprintf("%dpx", padding))
		})
	}
}

func PaddingT(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Padding)
		}

		listenBreakpointOption(widget.Padding, func(padding int) {
			widget.UpdateStyleProperty("padding-top", fmt.Sprintf("%dpx", padding))
		})
	}
}

func PaddingV(options ...breakpoint.BreakpointOptions[int]) Option {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Padding)
		}

		listenBreakpointOption(widget.Padding, func(padding int) {
			widget.UpdateStyleProperty("padding-top", fmt.Sprintf("%dpx", padding))
			widget.UpdateStyleProperty("padding-bottom", fmt.Sprintf("%dpx", padding))
		})
	}
}

func RowGap(rowGap int) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("row-gap", fmt.Sprintf("%dpx", rowGap))
	}
}

func SetText(text string) Option {
	return func(widget widget.BaseWidget) {
		widget.SetText(text)
	}
}

func Tooltip(tooltip string) Option {
	return func(widget widget.BaseWidget) {
		widget.SetAttribute("title", tooltip)
	}
}

func Transition(seconds float64) Option {
	return func(widget widget.BaseWidget) {
		if seconds < 0 {
			seconds = 0.3
		}

		widget.UpdateStyleProperty("transition", fmt.Sprintf("%.1fs", seconds))
	}
}

func UserSelect(userSelect UserSelectType) Option {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("user-select", string(userSelect))
	}
}

func Visible(options ...breakpoint.BreakpointOptions[bool]) Option {
	return func(widget widget.BaseWidget) {
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
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Width)
		}

		listenBreakpointOption(widget.Width, func(width int) {
			widget.UpdateStyleProperty("width", fmt.Sprintf("%dpx", width))
		})
	}
}

func WidthP(options ...breakpoint.BreakpointOptions[float64]) Option {
	return func(widget widget.BaseWidget) {
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
