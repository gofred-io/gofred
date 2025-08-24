package main

import (
	"github.com/man.go/mango/center"
	"github.com/man.go/mango/console"
	"github.com/man.go/mango/div"
	"github.com/man.go/mango/style"
	"github.com/man.go/mango/text"
	"github.com/man.go/mango/widget"
)

func main() {
	ctx := widget.Context()

	ctx.AppendChild(
		div.New(
			ctx,
			div.Style(
				div.Display(style.Display{Display: style.DisplayTypeFlex, FlexDirection: "row", Flex: 1}),
			),
			div.Children(
				div.New(
					ctx,
					div.Style(
						div.Display(style.Display{Flex: 1}),
						div.Border(style.Border{Width: 1, Style: "solid", Color: "red"}),
						div.Background(style.Background{Color: "blue"}),
					),
					div.OnClick(func(this widget.Widget) {
						if ctx.GetElementByID("hello2") == nil {
							console.Error("hello2 not found")
						}
					}),
				),
				div.New(
					ctx,
					div.Style(
						div.Display(style.Display{Display: style.DisplayTypeFlex, FlexDirection: "column", Flex: 1}),
						div.Border(style.Border{Width: 1, Style: "solid", Color: "red"}),
						div.Background(style.Background{Color: "green"}),
					),
					div.Children(
						center.New(
							ctx,
							div.ID("test"),
							div.Children(
								text.New(
									ctx,
									text.ID("hello"),
									text.Text("Hello, World!"),
									text.Style(
										text.Font(
											text.Color("white"),
											text.Family("Arial"),
											text.Size(12),
											text.Weight("lighter"),
										),
									),
								),
							),
						),
					),
				),
			),
		),
	)

	ctx.AppendChild(
		div.New(
			ctx,
			div.Style(
				div.Display(style.Display{Flex: 1}),
			),
		),
	)

	select {}
}
