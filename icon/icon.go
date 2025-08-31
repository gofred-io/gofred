package icon

import (
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/path"
	"github.com/gofred-io/gofred/svg"
	"github.com/gofred-io/gofred/widget"
)

type icon struct {
	opts []svg.Option
}

func New(data icondata.IconData, opts ...Option) widget.BaseWidget {
	i := &icon{}

	opts = append(
		opts,
		Class("gf-icon"),
	)

	for _, option := range opts {
		option(i)
	}

	return svg.New(
		[]widget.BaseWidget{
			path.New(
				string(data),
			),
		},
		i.opts...,
	)
}
