package grid

import (
	"github.com/gofred-io/gofred/div"
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/widget"
)

type Grid struct {
	ColumnCount int
	ColumnGap   int
	RowGap      int
}

func New(children []widget.Widget, options ...Options) widget.Widget {
	grid := &Grid{}

	for _, option := range options {
		option(grid)
	}

	return div.New(
		children,
		div.Style(
			div.Display(style.Display{
				Display:   style.DisplayTypeGrid,
				ColumnGap: grid.ColumnGap,
				RowGap:    grid.RowGap,
			}),
			div.ColumnCount(grid.ColumnCount),
		),
	)
}
