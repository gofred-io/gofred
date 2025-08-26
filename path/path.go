package path

import "github.com/gofred-io/gofred/widget"

type Path struct {
	widget.Widget
}

func New(options ...Options) widget.Widget {
	path := &Path{
		Widget: widget.Context().CreateElementNS("http://www.w3.org/2000/svg", "path"),
	}

	for _, option := range options {
		option(path)
	}

	return path.Widget
}
