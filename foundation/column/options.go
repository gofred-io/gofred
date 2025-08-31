package column

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/options"
)

type Option func(center *column)

func display(display options.DisplayType) Option {
	return func(column *column) {
		column.opts = append(column.opts, div.Display(display))
	}
}

func flexDirection(flexDirection options.FlexDirectionType) Option {
	return func(column *column) {
		column.opts = append(column.opts, div.FlexDirection(flexDirection))
	}
}

func Class(class string) Option {
	return func(column *column) {
		column.opts = append(column.opts, div.Class(class))
	}
}

func ID(id string) Option {
	return func(column *column) {
		column.opts = append(column.opts, div.ID(id))
	}
}

func CrossAxisAlignment(crossAxisAlignment options.AxisAlignmentType) Option {
	return func(column *column) {
		column.opts = append(column.opts, div.JustifyContent(crossAxisAlignment))
	}
}

func Flex(flex int) Option {
	return func(column *column) {
		column.opts = append(column.opts, div.Flex(flex))
	}
}

func Gap(gap int) Option {
	return func(column *column) {
		column.opts = append(column.opts, div.RowGap(gap))
	}
}

func MainAxisAlignment(mainAxisAlignment options.AxisAlignmentType) Option {
	return func(column *column) {
		column.opts = append(column.opts, div.AlignItems(mainAxisAlignment))
	}
}
