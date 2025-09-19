package grid

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/theme"
)

type Option func(grid *grid)

func display(display theme.DisplayType) Option {
	return func(grid *grid) {
		grid.opts = append(grid.opts, div.Display(display))
	}
}

func Class(class string) Option {
	return func(grid *grid) {
		grid.opts = append(grid.opts, div.Class(class))
	}
}

func ColumnCount(opts ...breakpoint.BreakpointOptions[int]) Option {
	return func(grid *grid) {
		grid.opts = append(grid.opts, div.ColumnCount(opts...))
	}
}

func ColumnGap(columnGap int) Option {
	return func(grid *grid) {
		grid.opts = append(grid.opts, div.ColumnGap(columnGap))
	}
}

func ID(id string) Option {
	return func(grid *grid) {
		grid.opts = append(grid.opts, div.ID(id))
	}
}

func RowGap(rowGap int) Option {
	return func(grid *grid) {
		grid.opts = append(grid.opts, div.RowGap(rowGap))
	}
}
