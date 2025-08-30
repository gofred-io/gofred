package grid

type Options func(grid *Grid)

func ColumnCount(columnCount int) Options {
	return func(grid *Grid) {
		grid.ColumnCount = columnCount
	}
}

func ColumnGap(columnGap int) Options {
	return func(grid *Grid) {
		grid.ColumnGap = columnGap
	}
}

func RowGap(rowGap int) Options {
	return func(grid *Grid) {
		grid.RowGap = rowGap
	}
}
