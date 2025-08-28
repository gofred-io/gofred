package main

import (
	"fmt"

	"github.com/gofred-io/gofred/container"
	"github.com/gofred-io/gofred/icon"
	"github.com/gofred-io/gofred/icondata"
	"github.com/gofred-io/gofred/path"
	"github.com/gofred-io/gofred/row"
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/svg"
	"github.com/gofred-io/gofred/text"
	"github.com/gofred-io/gofred/widget"
)

func main() {
	widget.Context().AppendChild(
		container.New(
			row.New(
				[]widget.Widget{
					icon.New(
						icon.Data(icondata.HamburgerMenu),
					),
					youtubeIcon(),
					counter(),
				},
				row.MainAxisAlignment(style.JustifyContentTypeStart),
				row.CrossAxisAlignment(style.AlignItemsTypeCenter),
			),
			container.Style(
				container.Padding(style.Padding{
					Right: 16,
					Left:  16,
				}),
				container.Height(56),
			),
			container.OnClick(func(widget widget.Widget) {
				fmt.Println("clicked")
			}),
		),
	)

	select {}
}

func youtubeIcon() widget.Widget {
	return svg.New(
		[]widget.Widget{
			path.New(
				path.Data("M14.4848 20C14.4848 20 23.5695 20 25.8229 19.4C27.0917 19.06 28.0459 18.08 28.3808 16.87C29 14.65 29 9.98 29 9.98C29 9.98 29 5.34 28.3808 3.14C28.0459 1.9 27.0917 0.94 25.8229 0.61C23.5695 0 14.4848 0 14.4848 0C14.4848 0 5.42037 0 3.17711 0.61C1.9286 0.94 0.954148 1.9 0.59888 3.14C0 5.34 0 9.98 0 9.98C0 9.98 0 14.65 0.59888 16.87C0.954148 18.08 1.9286 19.06 3.17711 19.4C5.42037 20 14.4848 20 14.4848 20Z"),
				path.Fill("#FF0033"),
			),
			path.New(
				path.Data("M19 10L11.5 5.75V14.25L19 10Z"),
				path.Fill("#FFFFFF"),
			),
		},
		svg.Style(
			svg.Height(20),
			svg.Width(32),
		),
	)
}

func counter() widget.Widget {
	return text.New(
		text.Text("0"),
		text.Style(
			text.Font(text.Size(16), text.Color("#000000")),
		),
	)
}
