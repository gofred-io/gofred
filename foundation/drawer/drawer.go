package drawer

import (
	"fmt"
	"time"

	. "github.com/gofred-io/gofred/basic/div"
	. "github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/widget"
)

type FDrawer struct {
	*widget.BaseWidget
	barrier    *BDiv
	menu       *BDiv
	transition float64
}

func Drawer(child widget.Widget) *FDrawer {
	d := &FDrawer{
		BaseWidget: widget.New("div"),
		menu:       Div(nil).Class("gf-drawer-menu"),
		barrier:    Div(nil).Class("gf-drawer-barrier").Width(AllBP(0)),
	}

	d.Transition(0.5).Class("gf-drawer-container")

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

func (d *FDrawer) Show() {
	breakpoint := GetCurrentBreakPoint()
	d.barrier.GetBaseWidget().UpdateStyleProperty("width", "100%")

	if d.menu.GetBaseWidget().BVWidth != nil {
		width := d.menu.GetBaseWidget().BVWidth.Get(breakpoint)
		d.menu.GetBaseWidget().UpdateStyleProperty("width", fmt.Sprintf("%dpx", width))
	} else if d.menu.GetBaseWidget().BVWidthP != nil {
		widthP := d.menu.GetBaseWidget().BVWidthP.Get(breakpoint)
		d.menu.GetBaseWidget().UpdateStyleProperty("width", fmt.Sprintf("%f%%", widthP*100))
	}
}

func (d *FDrawer) Hide() {
	d.menu.GetBaseWidget().UpdateStyleProperty("width", "0")
	time.AfterFunc(time.Duration(d.transition)*time.Second, func() {
		d.barrier.GetBaseWidget().UpdateStyleProperty("width", "0")
	})
}
