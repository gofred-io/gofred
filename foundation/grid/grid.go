package grid

import (
	. "github.com/gofred-io/gofred/basic/div"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/widget"
)

type FGrid struct {
	*BDiv
}

func Grid(children []Widget) *FGrid {
	g := &FGrid{
		BDiv: Div(
			children,
		),
	}

	g.display(DisplayTypeGrid)
	return g
}
