package application

import "github.com/gofred-io/gofred/widget"

func Run(app widget.BaseWidget) {
	widget.Context().AppendChild(app)
	select {}
}
