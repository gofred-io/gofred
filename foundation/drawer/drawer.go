package drawer

import (
	"fmt"
	"time"

	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/container"
)

type Drawer struct {
	application.BaseWidget
	barrier    application.BaseWidget
	menu       application.BaseWidget
	opts       []div.Option
	transition float64

	builder DrawerBuilder
}

type DrawerBuilder func() application.BaseWidget

func New(builder DrawerBuilder, opts ...Option) *Drawer {
	d := &Drawer{
		BaseWidget: application.New("div"),
		menu:       container.New(application.Nil, container.Class("gf-drawer-menu")),
		barrier:    div.New(nil, div.Class("gf-drawer-barrier"), div.Width(breakpoint.All(0))),
		builder:    builder,
	}

	opts = append([]Option{
		Transition(0.5),
		Class("gf-drawer-container"),
	}, opts...)

	for _, option := range opts {
		option(d)
	}

	d.barrier.SetOnClick(func(this application.BaseWidget, e application.Event) {
		d.Hide()
	})

	d.menu.SetOnClick(func(this application.BaseWidget, e application.Event) {
		e.StopPropagation()
	})

	d.menu.UpdateStyleProperty("width", "0")
	d.menu.AppendChild(d.builder().Widget)

	d.Widget.AppendChild(d.menu.Widget)
	d.Widget.AppendChild(d.barrier.Widget)

	application.Context().Root.AppendChild(d.Widget)
	return d
}

func (d *Drawer) Show() {
	breakpoint := breakpoint.GetCurrent()
	d.barrier.UpdateStyleProperty("width", "100%")

	if d.menu.Width != nil {
		width := d.menu.Width.Get(breakpoint)
		d.menu.UpdateStyleProperty("width", fmt.Sprintf("%dpx", width))
	} else if d.menu.WidthP != nil {
		widthP := d.menu.WidthP.Get(breakpoint)
		d.menu.UpdateStyleProperty("width", fmt.Sprintf("%f%%", widthP*100))
	}
}

func (d *Drawer) Hide() {
	d.menu.UpdateStyleProperty("width", "0")
	time.AfterFunc(time.Duration(d.transition)*time.Second, func() {
		d.barrier.UpdateStyleProperty("width", "0")
	})
}

func (d *Drawer) Refresh() {
	d.menu.ChildAt(0).ReplaceWith(d.builder().Widget)
}
