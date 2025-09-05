package path

import (
	"github.com/gofred-io/gofred/widget"
)

type path struct {
	*widget.BaseWidget
}

func New(data string, opts ...Option) *path {
	path := &path{
		BaseWidget: widget.NewNS("http://www.w3.org/2000/svg", "path"),
	}

	opts = append(
		opts,
		d(data),
	)

	for _, option := range opts {
		option()(path)
	}

	return path
}
