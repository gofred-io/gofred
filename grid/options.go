package grid

import (
	"github.com/gofred-io/gofred/style/breakpoint"
)

type Options func(grid *Grid)

func ColumnCount(options ...breakpoint.BreakpointOptions[int]) Options {
	return func(grid *Grid) {
		if grid.columnCount == nil {
			grid.columnCount = &breakpoint.BreakpointValue[int]{}
		}

		for _, option := range options {
			option(grid.columnCount)
		}
	}
}

func ColumnGap(columnGap int) Options {
	return func(grid *Grid) {
		grid.columnGap = columnGap
	}
}

func RowGap(rowGap int) Options {
	return func(grid *Grid) {
		grid.rowGap = rowGap
	}
}
