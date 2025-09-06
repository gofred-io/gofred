package column

import (
	. "github.com/gofred-io/gofred/options"
)

func (column *FColumn) display(display DisplayType) *FColumn {
	column.BDiv.Display(display)
	return column
}

func (column *FColumn) flexDirection(flexDirection FlexDirectionType) *FColumn {
	column.BDiv.FlexDirection(flexDirection)
	return column
}

func (column *FColumn) Class(class string) *FColumn {
	column.BDiv.Class(class)
	return column
}

func (column *FColumn) ID(id string) *FColumn {
	column.BDiv.ID(id)
	return column
}

func (column *FColumn) CrossAxisAlignment(crossAxisAlignment AxisAlignmentType) *FColumn {
	column.BDiv.AlignItems(crossAxisAlignment)
	return column
}

func (column *FColumn) Flex(flex int) *FColumn {
	column.BDiv.Flex(flex)
	return column
}

func (column *FColumn) Gap(gap int) *FColumn {
	column.BDiv.RowGap(gap)
	return column
}

func (column *FColumn) MainAxisAlignment(mainAxisAlignment AxisAlignmentType) *FColumn {
	column.BDiv.JustifyContent(mainAxisAlignment)
	return column
}

func (column *FColumn) MainAxisSize(mainAxisSize AxisSizeType) *FColumn {
	column.BDiv.AlignSelf(mainAxisSize)
	return column
}
