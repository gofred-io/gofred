package icon

import (
	"github.com/gofred-io/gofred/basic/path"
	"github.com/gofred-io/gofred/basic/svg"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	"github.com/gofred-io/gofred/widget"
)

type icon struct {
	*widget.BaseWidget
	opts []svg.Option
}

func New(data icondata.IconData, opts ...Option) *icon {
	i := &icon{
		BaseWidget: &widget.BaseWidget{},
	}

	defaultOpts := []Option{
		Class("gf-icon"),
	}

	opts = append(
		defaultOpts,
		opts...,
	)

	for _, option := range opts {
		option(i)
	}

	i.Extend(
		svg.New(
			[]widget.Widget{
				path.New(
					string(data),
				),
			},
			i.opts...,
		),
	)
	return i
}
