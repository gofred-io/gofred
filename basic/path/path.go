package path

import (
	"github.com/gofred-io/gofred/widget"
)

type Path struct {
	widget.BaseWidget
}

func New(data string, opts ...Option) widget.BaseWidget {
	path := &Path{
		BaseWidget: widget.NewNS("http://www.w3.org/2000/svg", "path"),
	}

	opts = append(
		opts,
		d(data),
	)

	for _, option := range opts {
		option()(path.BaseWidget)
	}

	return path.BaseWidget
}
