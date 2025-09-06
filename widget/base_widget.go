package widget

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/options/spacing"
)

type BaseWidget struct {
	*JSWidget
	BVColumnCount *breakpoint.BreakpointValue[int]
	BVHeight      *breakpoint.BreakpointValue[int]
	BVMargin      *breakpoint.BreakpointValue[spacing.Spacing]
	BVMaxWidth    *breakpoint.BreakpointValue[int]
	BVPadding     *breakpoint.BreakpointValue[spacing.Spacing]
	BVVisible     *breakpoint.BreakpointValue[bool]
	BVWidth       *breakpoint.BreakpointValue[int]
	BVWidthP      *breakpoint.BreakpointValue[float64]
}

func newBaseWidget(widget *JSWidget) *BaseWidget {
	return &BaseWidget{
		JSWidget:      widget,
		BVColumnCount: &breakpoint.BreakpointValue[int]{},
		BVHeight:      &breakpoint.BreakpointValue[int]{},
		BVMargin:      &breakpoint.BreakpointValue[spacing.Spacing]{},
		BVMaxWidth:    &breakpoint.BreakpointValue[int]{},
		BVPadding:     &breakpoint.BreakpointValue[spacing.Spacing]{},
		BVVisible:     &breakpoint.BreakpointValue[bool]{},
		BVWidth:       &breakpoint.BreakpointValue[int]{},
		BVWidthP:      &breakpoint.BreakpointValue[float64]{},
	}
}

func New(htmlTag string) *BaseWidget {
	return newBaseWidget(Context().CreateElement(htmlTag))
}

func NewWithOptions(htmlTag string, options map[string]any) *BaseWidget {
	return newBaseWidget(Context().CreateElementWithOptions(htmlTag, options))
}

func NewNS(namespace string, htmlTag string) *BaseWidget {
	return newBaseWidget(Context().CreateElementNS(namespace, htmlTag))
}

func (b *BaseWidget) GetBaseWidget() *BaseWidget {
	return b
}

func (b *BaseWidget) Extend(widget Widget) *BaseWidget {
	b.JSWidget = widget.GetBaseWidget().JSWidget
	return b
}

func (b *BaseWidget) ReplaceWith(widget Widget) {
	b.JSWidget.ReplaceWith(widget)
}
