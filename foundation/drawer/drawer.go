package drawer

import (
	"fmt"
	"time"

	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/widget"
)

type Drawer struct {
	*widget.BaseWidget
	barrier    widget.Widget
	menu       widget.Widget
	opts       []div.Option
	transition float64
}

func New(child widget.Widget, opts ...Option) *Drawer {
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

	d.barrier.GetBaseWidget().SetOnClick(func(this widget.Widget, e widget.Event) {
		d.Hide()
	})

	d.menu.GetBaseWidget().SetOnClick(func(this widget.Widget, e widget.Event) {
		e.StopPropagation()
	})

	d.menu.GetBaseWidget().UpdateStyleProperty("width", "0")
	d.menu.GetBaseWidget().AppendChild(child)

	d.JSWidget.AppendChild(d.menu)
	d.JSWidget.AppendChild(d.barrier)

	widget.Context().Root.AppendChild(d)
	return d
}

func (d *Drawer) Show() {
	breakpoint := breakpoint.GetCurrent()
	d.barrier.GetBaseWidget().UpdateStyleProperty("width", "100%")

	if d.menu.GetBaseWidget().Width != nil {
		width := d.menu.GetBaseWidget().Width.Get(breakpoint)
		d.menu.GetBaseWidget().UpdateStyleProperty("width", fmt.Sprintf("%dpx", width))
	} else if d.menu.GetBaseWidget().WidthP != nil {
		widthP := d.menu.GetBaseWidget().WidthP.Get(breakpoint)
		d.menu.GetBaseWidget().UpdateStyleProperty("width", fmt.Sprintf("%f%%", widthP*100))
	}
}

func (d *Drawer) Hide() {
	d.menu.GetBaseWidget().UpdateStyleProperty("width", "0")
	time.AfterFunc(time.Duration(d.transition)*time.Second, func() {
		d.barrier.GetBaseWidget().UpdateStyleProperty("width", "0")
	})
}
