package main

import (
	"fmt"

	"github.com/man.go/mango/div"
	"github.com/man.go/mango/style"
	"github.com/man.go/mango/widget"
)

func main() {
	ctx := widget.Context()
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
						fmt.Println("clicked")
					}),
				),
				div.New(
					ctx,
					div.Style(
						div.Size(32, 32),
						div.Border(style.Border{Width: 1, Style: "solid", Color: "red"}),
						div.Background(style.Background{Color: "green"}),
					),
				),
			),
		),
	)

	select {}
}
