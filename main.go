package main

import (
	"fmt"

	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/foundation/button"
	. "github.com/gofred-io/gofred/foundation/container"
	. "github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	. "github.com/gofred-io/gofred/foundation/row"
	. "github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/hooks"
	. "github.com/gofred-io/gofred/listenable"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/options/spacing"
	. "github.com/gofred-io/gofred/widget"
)

var (
	counter, setCounter = hooks.UseState(0)
)

func main() {
	Context().AppendChild(
		Container(
			Row(
				[]Widget{
					Icon(icondata.Menu),
					Button(Text("Count")).OnClick(func(widget Widget, e Event) {
						setCounter(counter.Value() + 1)
					}),
					counterWidget(),
				},
			).MainAxisAlignment(AxisAlignmentTypeStart).CrossAxisAlignment(AxisAlignmentTypeCenter),
		).Padding(AllBP(All(16))).Height(AllBP(56)),
	)

	select {}
}

func counterWidget() Widget {
	return Builder(counter, func() Widget {
		return Text(
			fmt.Sprintf("%d", counter.Value()),
		).FontSize(16).FontColor("#000000")
	})
}
