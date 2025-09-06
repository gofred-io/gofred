package row

import (
	. "github.com/gofred-io/gofred/options"
)

func (row *FRow) display(display DisplayType) *FRow {
	row.BDiv.Display(display)
	return row
}

func (row *FRow) flexDirection(flexDirection FlexDirectionType) *FRow {
	row.BDiv.FlexDirection(flexDirection)
	return row
}

func (row *FRow) Class(class string) *FRow {
	row.BDiv.Class(class)
	return row
}

func (row *FRow) ID(id string) *FRow {
	row.BDiv.ID(id)
	return row
}

func (row *FRow) CrossAxisAlignment(crossAxisAlignment AxisAlignmentType) *FRow {
	row.BDiv.AlignItems(crossAxisAlignment)
	return row
}

func (row *FRow) Flex(flex int) *FRow {
	row.BDiv.Flex(flex)
	return row
}

func (row *FRow) Gap(gap int) *FRow {
	row.BDiv.ColumnGap(gap)
	return row
}

func (row *FRow) MainAxisAlignment(mainAxisAlignment AxisAlignmentType) *FRow {
	row.BDiv.JustifyContent(mainAxisAlignment)
	return row
}

func (row *FRow) MainAxisSize(mainAxisSize AxisSizeType) *FRow {
	row.BDiv.AlignSelf(mainAxisSize)
	return row
}
