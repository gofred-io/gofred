package column

import (
	"fmt"

	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

type Options options.OptionsWrapper

func CrossAxisAlignment(crossAxisAlignment options.AxisAlignmentType) Options {
	return func() options.Options {
		return func(widget widget.BaseWidget) {
			widget.UpdateStyleProperty("justify-content", string(crossAxisAlignment))
		}
	}
}

func Gap(gap int) Options {
	return func() options.Options {
		return func(widget widget.BaseWidget) {
			widget.UpdateStyleProperty("row-gap", fmt.Sprintf("%dpx", gap))
		}
	}
}

func MainAxisAlignment(mainAxisAlignment options.AxisAlignmentType) Options {
	return func() options.Options {
		return func(widget widget.BaseWidget) {
			widget.UpdateStyleProperty("align-items", string(mainAxisAlignment))
		}
	}
}
