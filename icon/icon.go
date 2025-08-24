package icon

import (
	"github.com/man.go/mango/icondata"
	"github.com/man.go/mango/widget"
)

type Icon struct {
	widget.Widget
	Data icondata.IconData
}

func New(ctx *widget.WidgetContext, options ...Options) widget.Widget {
	icon := &Icon{
		Widget: ctx.CreateElement("span"),
	}

	for _, option := range options {
		option(icon)
	}

	icon.SetStyle("display: inline-block; width: 24px; height: 24px;")

	svgEl := ctx.CreateElementNS("http://www.w3.org/2000/svg", "svg")
	svgEl.SetAttribute("xmlns", "http://www.w3.org/2000/svg")
	svgEl.SetAttribute("viewBox", "0 0 24 24")
	svgEl.SetAttribute("width", "24")
	svgEl.SetAttribute("height", "24")
	svgEl.SetAttribute("focusable", "false")
	svgEl.SetAttribute("aria-hidden", "true")
	svgEl.SetStyle("pointer-events: none; display: inherit; width: 100%; height: 100%;")

	pathEl := ctx.CreateElementNS("http://www.w3.org/2000/svg", "path")
	pathEl.SetAttribute("d", string(icon.Data))
	svgEl.AppendChild(pathEl)

	icon.AppendChild(svgEl)

	return icon.Widget
}
