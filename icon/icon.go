package icon

import (
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/path"
	"github.com/gofred-io/gofred/svg"
	"github.com/gofred-io/gofred/widget"
)

func New(data icondata.IconData, opts ...options.Options) widget.BaseWidget {
	opts = append(
		opts,
		options.Class("gf-icon"),
	)

	return svg.New(
		[]widget.BaseWidget{
			path.New(
				options.D(string(data)),
			),
		},
		opts...,
	)
}
