package path

import "github.com/gofred-io/gofred/widget"

type Path struct {
	widget.Widget
}

func New(ctx *widget.WidgetContext, options ...Options) widget.Widget {
	path := &Path{
		Widget: ctx.CreateElementNS("http://www.w3.org/2000/svg", "path"),
	}

	for _, option := range options {
		option(path)
	}

	return path.Widget
}
