package main

import (
	"github.com/man.go/mango/div"
	"github.com/man.go/mango/style"
	"github.com/man.go/mango/text"
	"github.com/man.go/mango/widget"
)

func main() {
	ctx := widget.Context()
	div2 := buildDiv2(ctx)

	ctx.AppendChild(
		div.New(
			ctx,
			div.Style(
				div.Size(100, 100),
				div.Display(style.Display{Display: style.DisplayTypeFlex, FlexDirection: "row"}),
			),
			div.Children(
				div.New(
					ctx,
					div.Style(
						div.Size(32, 32),
						div.Border(style.Border{Width: 1, Style: "solid", Color: "red"}),
						div.Background(style.Background{Color: "blue"}),
					),
					div.OnClick(func(this widget.Widget) {
						ctx.GetElementByID("hello").SetText("Hello, again!")
					}),
				),
				div2,
			),
		),
	)

	select {}
}

func buildDiv2(ctx *widget.WidgetContext) widget.Widget {
	return div.New(
		ctx,
		div.ID("test"),
		div.Style(
			div.Size(32, 32),
			div.Border(style.Border{Width: 1, Style: "solid", Color: "red"}),
			div.Background(style.Background{Color: "green"}),
		),
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
	)
}
