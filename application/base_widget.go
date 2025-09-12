package application

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/options/spacing"
)

type BaseWidget struct {
	Widget
	ColumnCount *breakpoint.BreakpointValue[int]
	Height      *breakpoint.BreakpointValue[int]
	Margin      *breakpoint.BreakpointValue[spacing.Spacing]
	MaxWidth    *breakpoint.BreakpointValue[int]
	Padding     *breakpoint.BreakpointValue[spacing.Spacing]
	Visible     *breakpoint.BreakpointValue[bool]
	Width       *breakpoint.BreakpointValue[int]
	WidthP      *breakpoint.BreakpointValue[float64]
}

func newBaseWidget(widget Widget) BaseWidget {
	return BaseWidget{
		Widget:      widget,
		ColumnCount: &breakpoint.BreakpointValue[int]{},
		Height:      &breakpoint.BreakpointValue[int]{},
		Margin:      &breakpoint.BreakpointValue[spacing.Spacing]{},
		MaxWidth:    &breakpoint.BreakpointValue[int]{},
		Padding:     &breakpoint.BreakpointValue[spacing.Spacing]{},
		Visible:     &breakpoint.BreakpointValue[bool]{},
		Width:       &breakpoint.BreakpointValue[int]{},
		WidthP:      &breakpoint.BreakpointValue[float64]{},
	}
}

func New(htmlTag string) BaseWidget {
	return newBaseWidget(Context().CreateElement(htmlTag))
}

func NewWithOptions(htmlTag string, options map[string]any) BaseWidget {
	return newBaseWidget(Context().CreateElementWithOptions(htmlTag, options))

}

func NewNS(namespace string, htmlTag string) BaseWidget {
	return newBaseWidget(Context().CreateElementNS(namespace, htmlTag))

}
