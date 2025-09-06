package application

import (
	. "github.com/gofred-io/gofred/widget"
)

func Run(app Widget) {
	Context().AppendChild(app)
	select {}
}
