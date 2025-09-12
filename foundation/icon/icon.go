package icon

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/basic/path"
	"github.com/gofred-io/gofred/basic/svg"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
)

type icon struct {
	opts []svg.Option
}

func New(data icondata.IconData, opts ...Option) application.BaseWidget {
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
		[]application.BaseWidget{
			path.New(
				string(data),
			),
		},
		i.opts...,
	)
}
