package main

import (
	"log"

	"github.com/man.go/mango/div"
	"github.com/man.go/mango/icon"
	"github.com/man.go/mango/icondata"
	"github.com/man.go/mango/style"
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
							Display:       "flex",
							FlexDirection: style.FlexDirectionTypeRow,
							Flex:          1,
						}),
						div.Size(0, 56),
					),
					div.OnClick(func(this widget.Widget) {
						log.Printf("clicked")
						//ctx.GetElementByID("qwe0").SetAttribute("style", "background-color: red;")
					}),
					div.Children(
						icon.New(
							ctx,
							icon.Data(icondata.IconDataHamburgerMenu),
						),
					),
				),
			),
		),
	)

	select {}
}
