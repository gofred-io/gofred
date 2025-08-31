package style

import (
	"fmt"
	"strconv"

	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/style/breakpoint"
	"github.com/gofred-io/gofred/widget"
)

type Options func(widget widget.BaseWidget)
type OnClickHandler func(this widget.BaseWidget)

func AlignContent(alignContent AlignContentType) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("align-content", string(alignContent))
	}
}

func AlignItems(alignItems AlignItemsType) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("align-items", string(alignItems))
	}
}

func AlignSelf(alignSelf AlignSelfType) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("align-self", string(alignSelf))
	}
}

func Alt(alt string) Options {
	return func(widget widget.BaseWidget) {
		widget.SetAttribute("alt", alt)
	}
}

func Class(class string) Options {
	return func(widget widget.BaseWidget) {
		widget.AddClass(class)
	}
}

func FontColor(color string) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("color", color)
	}
}

func ColumnCount(columnCount ...breakpoint.BreakpointOptions[int]) Options {
	return func(widget widget.BaseWidget) {
		for _, option := range columnCount {
			option(widget.ColumnCount)
		}

		listenBreakpointOption(widget.ColumnCount, func(columnCount int) {
			widget.UpdateStyleProperty("grid-template-columns", fmt.Sprintf("repeat(%d, auto)", columnCount))
		})
	}
}

func ColumnGap(columnGap int) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("column-gap", fmt.Sprintf("%dpx", columnGap))
	}
}

func CrossAxisAlignment(crossAxisAlignment AxisAlignmentType) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("justify-content", string(crossAxisAlignment))
	}
}

func D(data string) Options {
	return func(widget widget.BaseWidget) {
		widget.SetAttribute("d", data)
	}
}

func Display(display DisplayType) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("display", string(display))
	}
}

func Fill(fill string) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("fill", fill)
	}
}

func Flex(flex int) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("flex", strconv.Itoa(flex))
	}
}

func FlexDirection(flexDirection FlexDirectionType) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("flex-direction", string(flexDirection))
	}
}

func FlexWrap(flexWrap FlexWrapType) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("flex-wrap", string(flexWrap))
	}
}

func FontFamily(fontFamily string) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("font-family", fontFamily)
	}
}

func FontSize(fontSize int) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("font-size", fmt.Sprintf("%dpx", fontSize))
	}
}

func FontWeight(fontWeight string) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("font-weight", fontWeight)
	}
}

func Height(options ...breakpoint.BreakpointOptions[int]) Options {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Height)

			listenBreakpointOption(widget.Height, func(height int) {
				widget.UpdateStyleProperty("height", fmt.Sprintf("%dpx", height))
			})
		}
	}
}

func Href(href string) Options {
	return func(widget widget.BaseWidget) {
		widget.SetAttribute("href", href)
	}
}

func ID(id string) Options {
	return func(widget widget.BaseWidget) {
		widget.SetID(id)
	}
}

func JustifyContent(justifyContent JustifyContentType) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("justify-content", string(justifyContent))
	}
}

func JustifyItems(justifyItems JustifyItemsType) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("justify-items", string(justifyItems))
	}
}

func JustifySelf(justifySelf JustifySelfType) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("justify-self", string(justifySelf))
	}
}

func LineHeight(lineHeight float64) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("line-height", fmt.Sprintf("%f", lineHeight))
	}
}

func MainAxisAlignment(mainAxisAlignment AxisAlignmentType) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("justify-content", string(mainAxisAlignment))
	}
}

func Margin(options ...breakpoint.BreakpointOptions[int]) Options {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Margin)
		}

		listenBreakpointOption(widget.Margin, func(margin int) {
			widget.UpdateStyleProperty("margin", fmt.Sprintf("%dpx", margin))
		})
	}
}

func MarginB(options ...breakpoint.BreakpointOptions[int]) Options {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Margin)
		}

		listenBreakpointOption(widget.Margin, func(margin int) {
			widget.UpdateStyleProperty("margin-bottom", fmt.Sprintf("%dpx", margin))
		})
	}
}

func MarginL(options ...breakpoint.BreakpointOptions[int]) Options {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Margin)
		}

		listenBreakpointOption(widget.Margin, func(margin int) {
			widget.UpdateStyleProperty("margin-left", fmt.Sprintf("%dpx", margin))
		})
	}
}

func MarginR(options ...breakpoint.BreakpointOptions[int]) Options {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Margin)
		}

		listenBreakpointOption(widget.Margin, func(margin int) {
			widget.UpdateStyleProperty("margin-right", fmt.Sprintf("%dpx", margin))
		})
	}
}

func MarginT(options ...breakpoint.BreakpointOptions[int]) Options {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Margin)
		}

		listenBreakpointOption(widget.Margin, func(margin int) {
			widget.UpdateStyleProperty("margin-top", fmt.Sprintf("%dpx", margin))
		})
	}
}

func OnClick(handler OnClickHandler) Options {
	return func(widget widget.BaseWidget) {
		widget.SetOnClick(handler)
	}
}

func NewTab(newTab bool) Options {
	return func(widget widget.BaseWidget) {
		if newTab {
			widget.SetAttribute("target", "_blank")
		}
	}
}

func Padding(options ...breakpoint.BreakpointOptions[int]) Options {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Padding)
		}

		listenBreakpointOption(widget.Padding, func(padding int) {
			widget.UpdateStyleProperty("padding", fmt.Sprintf("%dpx", padding))
		})
	}
}

func PaddingB(options ...breakpoint.BreakpointOptions[int]) Options {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Padding)
		}

		listenBreakpointOption(widget.Padding, func(padding int) {
			widget.UpdateStyleProperty("padding-bottom", fmt.Sprintf("%dpx", padding))
		})
	}
}

func PaddingL(options ...breakpoint.BreakpointOptions[int]) Options {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Padding)
		}

		listenBreakpointOption(widget.Padding, func(padding int) {
			widget.UpdateStyleProperty("padding-left", fmt.Sprintf("%dpx", padding))
		})
	}
}

func PaddingR(options ...breakpoint.BreakpointOptions[int]) Options {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Padding)
		}

		listenBreakpointOption(widget.Padding, func(padding int) {
			widget.UpdateStyleProperty("padding-right", fmt.Sprintf("%dpx", padding))
		})
	}
}

func PaddingT(options ...breakpoint.BreakpointOptions[int]) Options {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Padding)
		}

		listenBreakpointOption(widget.Padding, func(padding int) {
			widget.UpdateStyleProperty("padding-top", fmt.Sprintf("%dpx", padding))
		})
	}
}

func RowGap(rowGap int) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("row-gap", fmt.Sprintf("%dpx", rowGap))
	}
}

func Tooltip(tooltip string) Options {
	return func(widget widget.BaseWidget) {
		widget.SetAttribute("title", tooltip)
	}
}

func UserSelect(userSelect UserSelectType) Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("user-select", string(userSelect))
	}
}

func Visible(options ...breakpoint.BreakpointOptions[bool]) Options {
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

func Width(options ...breakpoint.BreakpointOptions[int]) Options {
	return func(widget widget.BaseWidget) {
		for _, option := range options {
			option(widget.Width)
		}

		listenBreakpointOption(widget.Width, func(width int) {
			widget.UpdateStyleProperty("width", fmt.Sprintf("%dpx", width))
		})
	}
}

func listenBreakpointOption[T any](breakpointValue *breakpoint.BreakpointValue[T], callback func(value T)) {
	breakPointValue := hooks.UseBreakpoint()
	innerCallback := func(breakPoint breakpoint.BreakPoint) {
		if breakpointValue == nil {
			return
		}

		value := breakpointValue.Get(breakPoint)
		callback(value)
	}

	listener := listenable.NewListener(innerCallback)
	breakPointValue.AddListener(listener)

	currentBreakpoint := breakpoint.GetCurrent()
	innerCallback(currentBreakpoint)
}
