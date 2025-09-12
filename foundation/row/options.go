package row

import (
	"github.com/gofred-io/gofred/basic/div"
	"github.com/gofred-io/gofred/theme"
)

type Option func(center *row)

func display(display theme.DisplayType) Option {
	return func(row *row) {
		row.opts = append(row.opts, div.Display(display))
	}
}

func flexDirection(flexDirection theme.FlexDirectionType) Option {
	return func(row *row) {
		row.opts = append(row.opts, div.FlexDirection(flexDirection))
	}
}

func Class(class string) Option {
	return func(row *row) {
		row.opts = append(row.opts, div.Class(class))
	}
}

func ID(id string) Option {
	return func(row *row) {
		row.opts = append(row.opts, div.ID(id))
	}
}

func CrossAxisAlignment(crossAxisAlignment theme.AxisAlignmentType) Option {
	return func(row *row) {
		row.opts = append(row.opts, div.AlignItems(crossAxisAlignment))
	}
}

func Flex(flex int) Option {
	return func(row *row) {
		row.opts = append(row.opts, div.Flex(flex))
	}
}

func Gap(gap int) Option {
	return func(row *row) {
		row.opts = append(row.opts, div.ColumnGap(gap))
	}
}

func MainAxisAlignment(mainAxisAlignment theme.AxisAlignmentType) Option {
	return func(row *row) {
		row.opts = append(row.opts, div.JustifyContent(mainAxisAlignment))
	}
}

func MainAxisSize(mainAxisSize theme.AxisSizeType) Option {
	return func(row *row) {
		row.opts = append(row.opts, div.AlignSelf(mainAxisSize))
	}
}
