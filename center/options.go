package center

import (
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

func alignItems(alignItems options.AxisAlignmentType) options.Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("align-items", string(alignItems))
	}
}

func justifyContent(justifyContent options.AxisAlignmentType) options.Options {
	return func(widget widget.BaseWidget) {
		widget.UpdateStyleProperty("justify-content", string(justifyContent))
	}
}
