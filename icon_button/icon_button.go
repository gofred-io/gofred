package iconbutton

import (
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/icon"
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/style/breakpoint"
	"github.com/gofred-io/gofred/widget"
)

type IconButton struct {
	widget.Widget
	style   style.Style
	tooltip string
	visible *breakpoint.BreakpointValue[bool]
}

func New(iconData icondata.IconData, options ...Options) widget.Widget {
	iconButton := &IconButton{
		Widget: widget.Context().CreateElement("button"),
	}

	iconButton.AddClass("gf-icon-button")

	for _, option := range options {
		option(iconButton)
	}

	if iconButton.tooltip != "" {
		iconButton.SetAttribute("title", iconButton.tooltip)
	}

	iconButton.SetStyle(iconButton.style.String())

	iconElement := icon.New(iconData)
	iconButton.AppendChild(iconElement)
	iconButton.visibleListener()

	return iconButton.Widget
}

func (i *IconButton) visibleListener() {
	breakPointValue := hooks.UseBreakpoint()
	callback := func(breakPoint breakpoint.BreakPoint) {
		if i.visible == nil {
			return
		}

		visible := i.visible.Get(breakPoint)
		if visible {
			i.RemoveClass("gf-hidden")
		} else {
			i.AddClass("gf-hidden")
		}
	}

	listener := listenable.NewListener(callback)
	breakPointValue.AddListener(listener)

	currentBreakpoint := breakpoint.GetCurrent()
	callback(currentBreakpoint)
}
