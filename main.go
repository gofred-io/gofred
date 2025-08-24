package main

import (
	"github.com/man.go/mango/div"
	"github.com/man.go/mango/icon"
	"github.com/man.go/mango/icondata"
	"github.com/man.go/mango/path"
	"github.com/man.go/mango/style"
	"github.com/man.go/mango/svg"
	"github.com/man.go/mango/widget"
)

func main() {
	ctx := widget.Context()

	ctx.AppendChild(
		div.New(
			ctx,
			div.Children(
				div.New(
					ctx,
					div.Style(
						div.Display(style.Display{
							Display:       style.DisplayTypeFlex,
							FlexDirection: style.FlexDirectionTypeRow,
							Flex:          1,
							AlignItems:    style.AlignItemsTypeCenter,
						}),
						div.Size(0, 56),
						div.Padding(style.Padding{
							Right: 16,
							Left:  16,
						}),
					),
					div.Children(
						icon.New(
							ctx,
							icon.Data(icondata.HamburgerMenu),
						),
						youtubeIcon(ctx),
					),
				),
			),
		),
	)

	select {}
}

func youtubeIcon(ctx *widget.WidgetContext) widget.Widget {
	return svg.New(
		ctx,
		svg.Height(20),
		svg.Width(32),
		svg.Children(
			path.New(
				ctx,
				path.Data("M14.4848 20C14.4848 20 23.5695 20 25.8229 19.4C27.0917 19.06 28.0459 18.08 28.3808 16.87C29 14.65 29 9.98 29 9.98C29 9.98 29 5.34 28.3808 3.14C28.0459 1.9 27.0917 0.94 25.8229 0.61C23.5695 0 14.4848 0 14.4848 0C14.4848 0 5.42037 0 3.17711 0.61C1.9286 0.94 0.954148 1.9 0.59888 3.14C0 5.34 0 9.98 0 9.98C0 9.98 0 14.65 0.59888 16.87C0.954148 18.08 1.9286 19.06 3.17711 19.4C5.42037 20 14.4848 20 14.4848 20Z"),
				path.Fill("#FF0033"),
			),
			path.New(
				ctx,
				path.Data("M19 10L11.5 5.75V14.25L19 10Z"),
				path.Fill("#FFFFFF"),
			),
		),
	)
}
