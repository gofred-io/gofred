package path

import (
	"github.com/gofred-io/gofred/application"
)

type Path struct {
	application.BaseWidget
}

func New(data string, opts ...Option) application.BaseWidget {
	path := &Path{
		BaseWidget: application.NewNS("http://www.w3.org/2000/svg", "path"),
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
