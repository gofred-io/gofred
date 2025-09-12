package iconbutton

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/basic/button"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
)

type iconButton struct {
	opts     []button.Option
	iconOpts []icon.Option
}

func New(iconData icondata.IconData, opts ...Option) application.BaseWidget {
	iconButton := &iconButton{}

	defaultOpts := []Option{
		Class("gf-icon-button"),
	}

	opts = append(defaultOpts, opts...)

	for _, option := range opts {
		option(iconButton)
	}

	iconElement := icon.New(iconData, iconButton.iconOpts...)
	return button.New(iconElement, iconButton.opts...)
}
