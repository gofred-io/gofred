package main

import (
	"fmt"

	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/button"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

var (
	counter, setCounter = hooks.UseState(0)
)

func main() {
	application.Run(
		container.New(
			row.New(
				[]application.BaseWidget{
					icon.New(
						icondata.Menu,
					),
					button.New(
						text.New(
							"Count",
						),
						button.OnClick(func(widget application.BaseWidget, e application.Event) {
							setCounter(counter.Value() + 1)
						}),
					),
					counterWidget(),
				},
				row.MainAxisAlignment(theme.AxisAlignmentTypeStart),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
			container.Padding(breakpoint.All(spacing.All(16))),
			container.Height(breakpoint.All(56)),
		),
	)
}

func counterWidget() application.BaseWidget {
	return listenable.Builder(counter, func() application.BaseWidget {
		return text.New(
			fmt.Sprintf("%d", counter.Value()),
			text.FontSize(16),
			text.FontColor("#000000"),
		)
	})
}
