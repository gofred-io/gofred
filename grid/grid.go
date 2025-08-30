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

type Grid struct {
	div.Div
	columnCount *breakpoint.BreakpointValue[int]
	columnGap   int
	rowGap      int
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

	breakPointValue := hooks.UseBreakpoint()
	listener := listenable.NewListener(func(breakPoint breakpoint.BreakPoint) {
		currentColumnCount := grid.GetColumnCount(breakPoint)
		if currentColumnCount != nil {
			grid.Get("style").Set("gridTemplateColumns", fmt.Sprintf("repeat(%d, auto)", *currentColumnCount))
		}
	})
	breakPointValue.AddListener(listener)

	return grid.Widget
}

func (g *Grid) GetColumnCount(breakPoint breakpoint.BreakPoint) *int {
	columnCount := g.columnCount.Get(breakPoint)
	if columnCount == 0 {
		return nil
	}

	return &columnCount
}
