package main

import (
	"fmt"

	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/container"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/icon"
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/path"
	"github.com/gofred-io/gofred/row"
	"github.com/gofred-io/gofred/svg"
	"github.com/gofred-io/gofred/text"
	"github.com/gofred-io/gofred/widget"
)

var (
	counter, setCounter = hooks.UseState(0)
)

func main() {
	widget.Context().AppendChild(
		container.New(
			row.New(
				[]widget.BaseWidget{
					icon.New(
						icondata.Menu,
					),
					youtubeIcon(),
					counterWidget(),
				},
				row.MainAxisAlignment(options.AxisAlignmentTypeStart),
				row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
			),
			options.Padding(breakpoint.All(16)),
			options.Height(breakpoint.All(56)),
			options.OnClick(func(widget widget.BaseWidget) {
				setCounter(counter.Value() + 1)
			}),
		),
	)

	select {}
}

func youtubeIcon() widget.BaseWidget {
	return svg.New(
		[]widget.BaseWidget{
			path.New(
				options.D("M14.4848 20C14.4848 20 23.5695 20 25.8229 19.4C27.0917 19.06 28.0459 18.08 28.3808 16.87C29 14.65 29 9.98 29 9.98C29 9.98 29 5.34 28.3808 3.14C28.0459 1.9 27.0917 0.94 25.8229 0.61C23.5695 0 14.4848 0 14.4848 0C14.4848 0 5.42037 0 3.17711 0.61C1.9286 0.94 0.954148 1.9 0.59888 3.14C0 5.34 0 9.98 0 9.98C0 9.98 0 14.65 0.59888 16.87C0.954148 18.08 1.9286 19.06 3.17711 19.4C5.42037 20 14.4848 20 14.4848 20Z"),
				options.Fill("#FF0033"),
			),
			path.New(
				options.D("M19 10L11.5 5.75V14.25L19 10Z"),
				options.Fill("#FFFFFF"),
			),
		},
		options.Height(breakpoint.All(20)),
		options.Width(breakpoint.All(32)),
	)
}

func counterWidget() widget.BaseWidget {
	return listenable.Builder(counter, func() widget.BaseWidget {
		return text.New(
			fmt.Sprintf("%d", counter.Value()),
			options.FontSize(16),
			options.FontColor("#000000"),
		)
	})
}
