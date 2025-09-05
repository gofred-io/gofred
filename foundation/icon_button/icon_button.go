package iconbutton

import (
	"github.com/gofred-io/gofred/basic/button"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	"github.com/gofred-io/gofred/widget"
)

type iconButton struct {
	*widget.BaseWidget
	opts     []button.Option
	iconOpts []icon.Option
}

func New(iconData icondata.IconData, opts ...Option) *iconButton {
	iconButton := &iconButton{
		BaseWidget: &widget.BaseWidget{},
	}

	defaultOpts := []Option{
		Class("gf-icon-button"),
	}

	opts = append(defaultOpts, opts...)

	for _, option := range opts {
		option(iconButton)
	}

	iconElement := icon.New(iconData, iconButton.iconOpts...)
	iconButton.Extend(
		button.New(iconElement, iconButton.opts...),
	)

	return iconButton
}
