package icon

import (
	"github.com/gofred-io/gofred/basic/path"
	"github.com/gofred-io/gofred/basic/svg"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	"github.com/gofred-io/gofred/widget"
)

type icon struct {
	opts []svg.Option
}

func New(data icondata.IconData, opts ...Option) widget.BaseWidget {
	i := &icon{}

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

	return svg.New(
		[]widget.BaseWidget{
			path.New(
				string(data),
			),
		},
		i.opts...,
	)
}
