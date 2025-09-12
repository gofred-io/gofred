package center

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/theme"
)

type Option func(center *center)

func alignItems(alignItems theme.AxisAlignmentType) Option {
	return func(center *center) {
		center.opts = append(center.opts, div.AlignItems(alignItems))
	}
}

func display(display theme.DisplayType) Option {
	return func(center *center) {
		center.opts = append(center.opts, div.Display(display))
	}
}

func flex(flex int) Option {
	return func(center *center) {
		center.opts = append(center.opts, div.Flex(flex))
	}
}

func justifyContent(justifyContent theme.AxisAlignmentType) Option {
	return func(center *center) {
		center.opts = append(center.opts, div.JustifyContent(justifyContent))
	}
}

func Class(class string) Option {
	return func(center *center) {
		center.opts = append(center.opts, div.Class(class))
	}
}

func ID(id string) Option {
	return func(center *center) {
		center.opts = append(center.opts, div.ID(id))
	}
}
