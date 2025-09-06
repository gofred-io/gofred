package container

import (
	. "github.com/gofred-io/gofred/basic/div"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/widget"
)

type FContainer struct {
	*BDiv
}

func Container(child Widget) *FContainer {
	var (
		c = &FContainer{
			BDiv: Div(
				[]Widget{child},
			),
		}
	)

	c.display(DisplayTypeFlex)
	return c
}
