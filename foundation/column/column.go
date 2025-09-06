package column

import (
	. "github.com/gofred-io/gofred/basic/div"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/widget"
)

type FColumn struct {
	*BDiv
}

func Column(children []Widget) *FColumn {
	c := &FColumn{
		BDiv: Div(
			children,
		),
	}

	c.display(DisplayTypeFlex)
	c.flexDirection(FlexDirectionTypeColumn)

	return c
}
