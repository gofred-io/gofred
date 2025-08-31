package widget

import "github.com/gofred-io/gofred/style/breakpoint"

type BaseWidget struct {
	Widget
	ColumnCount *breakpoint.BreakpointValue[int]
	Height      *breakpoint.BreakpointValue[int]
	Margin      *breakpoint.BreakpointValue[int]
	Padding     *breakpoint.BreakpointValue[int]
	Visible     *breakpoint.BreakpointValue[bool]
	Width       *breakpoint.BreakpointValue[int]
}

func newBaseWidget(widget Widget) BaseWidget {
	return BaseWidget{
		Widget:      widget,
		ColumnCount: &breakpoint.BreakpointValue[int]{},
		Height:      &breakpoint.BreakpointValue[int]{},
		Margin:      &breakpoint.BreakpointValue[int]{},
		Padding:     &breakpoint.BreakpointValue[int]{},
		Visible:     &breakpoint.BreakpointValue[bool]{},
		Width:       &breakpoint.BreakpointValue[int]{},
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
