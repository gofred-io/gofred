package column

import (
	"github.com/gofred-io/gofred/div"
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/widget"
)

type Column struct {
	div.Div
	style style.Style
}

func New(children []widget.Widget, options ...Options) widget.Widget {
	column := &Column{
		style: style.Style{
			Display: &style.Display{
				Display:       style.DisplayTypeFlex,
				FlexDirection: style.FlexDirectionTypeColumn,
				Flex:          1,
			},
		},
	}

	for _, option := range options {
		option(column)
	}

	return div.New(
		children,
		div.StyleFrom(&column.style),
	)
}
