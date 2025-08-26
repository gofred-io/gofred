package row

import (
	"github.com/gofred-io/gofred/div"
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/widget"
)

type Row struct {
	div.Div
	style style.Style
}

func New(children []widget.Widget, options ...Options) widget.Widget {
	row := &Row{
		style: style.Style{
			Display: &style.Display{
				Display:       style.DisplayTypeFlex,
				FlexDirection: style.FlexDirectionTypeRow,
				Flex:          1,
			},
		},
	}

	for _, option := range options {
		option(row)
	}

	return div.New(
		children,
		div.StyleFrom(&row.style),
	)
}
