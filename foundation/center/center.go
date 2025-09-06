package center

import (
	. "github.com/gofred-io/gofred/basic/div"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/widget"
)

type FCenter struct {
	*BDiv
}

func Center(child Widget) *FCenter {
	c := &FCenter{
		BDiv: Div(
			[]Widget{child},
		),
	}

	c.display(DisplayTypeFlex)
	c.flex(1)
	c.alignItems(AxisAlignmentTypeCenter)
	c.justifyContent(AxisAlignmentTypeCenter)

	return c
}
