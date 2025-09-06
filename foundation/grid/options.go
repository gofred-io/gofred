package grid

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/options"
)

func (grid *FGrid) display(display DisplayType) *FGrid {
	grid.BDiv.Display(display)
	return grid
}

func (grid *FGrid) Class(class string) *FGrid {
	grid.BDiv.Class(class)
	return grid
}

func (grid *FGrid) ColumnCount(opts ...BreakpointOptions[int]) *FGrid {
	grid.BDiv.ColumnCount(opts...)
	return grid
}

func (grid *FGrid) ColumnGap(columnGap int) *FGrid {
	grid.BDiv.ColumnGap(columnGap)
	return grid
}

func (grid *FGrid) ID(id string) *FGrid {
	grid.BDiv.ID(id)
	return grid
}

func (grid *FGrid) RowGap(rowGap int) *FGrid {
	grid.BDiv.RowGap(rowGap)
	return grid
}
