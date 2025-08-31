package main

import (
	"fmt"

	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/button"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/icon"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/hooks"
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/options"
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
					button.New(
						text.New(
							"Count",
						),
						button.OnClick(func(widget widget.BaseWidget) {
							setCounter(counter.Value() + 1)
						}),
					),
					counterWidget(),
				},
				row.MainAxisAlignment(options.AxisAlignmentTypeStart),
				row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
			),
			container.Padding(breakpoint.All(16)),
			container.Height(breakpoint.All(56)),
		),
	)

	select {}
}

func counterWidget() widget.BaseWidget {
	return listenable.Builder(counter, func() widget.BaseWidget {
		return text.New(
			fmt.Sprintf("%d", counter.Value()),
			text.FontSize(16),
			text.FontColor("#000000"),
		)
	})
}
