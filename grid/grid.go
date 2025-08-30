package grid

import (
	"fmt"

	"github.com/gofred-io/gofred/div"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/style/breakpoint"
	"github.com/gofred-io/gofred/widget"
)

type GridColumnCount []int

type Grid struct {
	div.Div
	columnCount *GridColumnCount
	columnGap   int
	rowGap      int

	// listeners
	breakPointUpdatedListener *hooks.BreakpointValue
}

func New(children []widget.Widget, options ...Options) widget.Widget {
	grid := &Grid{}

	for _, option := range options {
		option(grid)
	}

	divStyleOptions := []div.StyleOptions{
		div.Display(style.Display{
			Display:   style.DisplayTypeGrid,
			ColumnGap: grid.columnGap,
			RowGap:    grid.rowGap,
		}),
	}

	if grid.columnCount != nil {
		currentBreakpoint := breakpoint.GetCurrent()
		currentColumnCount := grid.GetColumnCount(currentBreakpoint)
		if currentColumnCount != nil {
			divStyleOptions = append(divStyleOptions, div.ColumnCount(*currentColumnCount))
		}
	}

	grid.Widget = div.New(
		children,
		div.ID("grid"),
		div.Style(
			divStyleOptions...,
		),
	)

	grid.breakPointUpdatedListener = hooks.UseBreakpoint()
	listener := listenable.NewListener(func(breakPoint breakpoint.BreakPoint) {
		currentColumnCount := grid.GetColumnCount(breakPoint)
		if currentColumnCount != nil {
			grid.Get("style").Set("gridTemplateColumns", fmt.Sprintf("repeat(%d, auto)", *currentColumnCount))
		}
	})
	grid.breakPointUpdatedListener.AddListener(listener)

	return grid.Widget
}

func (g *Grid) GetColumnCount(breakPoint breakpoint.BreakPoint) *int {
	for i := int(breakPoint); i >= 0; i-- {
		if (*g.columnCount)[i] > 0 {
			return &(*g.columnCount)[i]
		}
	}

	return nil
}
