package drawer

import (
	"fmt"
	"time"

	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/widget"
)

type Drawer struct {
	widget.BaseWidget
	barrier    widget.BaseWidget
	menu       widget.BaseWidget
	opts       []div.Option
	transition float64
}

func New(opts ...Option) *Drawer {
	d := &Drawer{
		BaseWidget: widget.New("div"),
		menu:       div.New(nil, div.Class("gf-drawer-menu")),
		barrier:    div.New(nil, div.Class("gf-drawer-barrier"), div.Width(breakpoint.All(0))),
	}

	opts = append([]Option{Transition(0.5)}, opts...)
	opts = append(opts, Class("gf-drawer-container"))

	for _, option := range opts {
		option(d)
	}

	d.barrier.SetOnClick(func(this widget.BaseWidget, e widget.Event) {
		d.Hide()
	})

	d.menu.SetOnClick(func(this widget.BaseWidget, e widget.Event) {
		e.PreventDefault()
		e.StopPropagation()
	})

	d.menu.UpdateStyleProperty("width", "0")

	d.Widget.AppendChild(d.menu.Widget)
	d.Widget.AppendChild(d.barrier.Widget)

	widget.Context().Root.AppendChild(d.Widget)
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
