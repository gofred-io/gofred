package grid

import (
	"github.com/gofred-io/gofred/style/breakpoint"
)

type Options func(grid *Grid)
type ColumnCountOptions func(columnCount *GridColumnCount)

func ColumnCount(options ...ColumnCountOptions) Options {
	return func(grid *Grid) {
		if grid.columnCount == nil {
			columnCount := GridColumnCount(make([]int, len(breakpoint.BreakPoints)))
			grid.columnCount = &columnCount
		}

		for _, option := range options {
			option(grid.columnCount)
		}
	}
}

func ColumnCountXS(columnCount int) ColumnCountOptions {
	return func(gridColumnCount *GridColumnCount) {
		(*gridColumnCount)[breakpoint.XS] = columnCount
	}
}

func ColumnCountSM(columnCount int) ColumnCountOptions {
	return func(gridColumnCount *GridColumnCount) {
		(*gridColumnCount)[breakpoint.SM] = columnCount
	}
}

func ColumnCountMD(columnCount int) ColumnCountOptions {
	return func(gridColumnCount *GridColumnCount) {
		(*gridColumnCount)[breakpoint.MD] = columnCount
	}
}

func ColumnCountLG(columnCount int) ColumnCountOptions {
	return func(gridColumnCount *GridColumnCount) {
		(*gridColumnCount)[breakpoint.LG] = columnCount
	}
}

func ColumnCountXL(columnCount int) ColumnCountOptions {
	return func(gridColumnCount *GridColumnCount) {
		(*gridColumnCount)[breakpoint.XL] = columnCount
	}
}

func ColumnCountXXL(columnCount int) ColumnCountOptions {
	return func(gridColumnCount *GridColumnCount) {
		(*gridColumnCount)[breakpoint.XXL] = columnCount
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
