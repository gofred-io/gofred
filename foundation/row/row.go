package row

import (
	. "github.com/gofred-io/gofred/basic/div"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/widget"
)

type FRow struct {
	*BDiv
}

func Row(children []Widget) *FRow {
	r := &FRow{
		BDiv: Div(
			children,
		),
	}

	r.display(DisplayTypeFlex)
	r.flexDirection(FlexDirectionTypeRow)
	return r
}
