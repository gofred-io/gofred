package application

import "github.com/gofred-io/gofred/widget"

func Run(app widget.Widget) {
	widget.Context().AppendChild(app)
	select {}
}
