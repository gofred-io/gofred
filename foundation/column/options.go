package column

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/theme"
)

type Option func(center *column)

func display(display theme.DisplayType) Option {
	return func(column *column) {
		column.opts = append(column.opts, div.Display(display))
	}
}

func flexDirection(flexDirection theme.FlexDirectionType) Option {
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

func CrossAxisAlignment(crossAxisAlignment theme.AxisAlignmentType) Option {
	return func(column *column) {
		column.opts = append(column.opts, div.AlignItems(crossAxisAlignment))
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

func MainAxisAlignment(mainAxisAlignment theme.AxisAlignmentType) Option {
	return func(column *column) {
		column.opts = append(column.opts, div.JustifyContent(mainAxisAlignment))
	}
}

func MainAxisSize(mainAxisSize theme.AxisSizeType) Option {
	return func(column *column) {
		column.opts = append(column.opts, div.AlignSelf(mainAxisSize))
	}
}

func Overflow(overflow theme.OverflowType) Option {
	return func(column *column) {
		column.opts = append(column.opts, div.Overflow(overflow))
	}
}
