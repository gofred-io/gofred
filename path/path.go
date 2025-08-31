package path

import (
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type Path struct {
	widget.BaseWidget
}

func New(opts ...options.Options) widget.BaseWidget {
	path := &Path{
		BaseWidget: widget.NewNS("http://www.w3.org/2000/svg", "path"),
	}

	for _, option := range opts {
		option(path.BaseWidget)
	}

	return path.BaseWidget
}
