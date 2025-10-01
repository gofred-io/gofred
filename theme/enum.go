package theme

type AxisAlignmentType string
type AxisSizeType string
type BorderStyleType string
type DisplayType string
type FlexDirectionType string
type FlexWrapType string
type OverflowType string
type PositionType string
type PositionStyleType string
type TextAlignType string
type UserSelectType string

const (
	AxisAlignmentTypeCenter   AxisAlignmentType = "center"
	AxisAlignmentTypeStart    AxisAlignmentType = "start"
	AxisAlignmentTypeEnd      AxisAlignmentType = "end"
	AxisAlignmentTypeStretch  AxisAlignmentType = "stretch"
	AxisAlignmentTypeBaseline AxisAlignmentType = "baseline"
)

const (
	AxisSizeTypeMax AxisSizeType = "stretch"
	AxisSizeTypeMin AxisSizeType = "start"
)

const (
	BorderStyleTypeSolid  BorderStyleType = "solid"
	BorderStyleTypeDashed BorderStyleType = "dashed"
	BorderStyleTypeDotted BorderStyleType = "dotted"
	BorderStyleTypeDouble BorderStyleType = "double"
	BorderStyleTypeGroove BorderStyleType = "groove"
	BorderStyleTypeRidge  BorderStyleType = "ridge"
)

const (
	DisplayTypeFlex DisplayType = "flex"
	DisplayTypeGrid DisplayType = "grid"
	DisplayTypeNone DisplayType = "none"
)

const (
	FlexDirectionTypeRow    FlexDirectionType = "row"
	FlexDirectionTypeColumn FlexDirectionType = "column"
)

const (
	FlexWrapTypeNowrap      FlexWrapType = "nowrap"
	FlexWrapTypeWrap        FlexWrapType = "wrap"
	FlexWrapTypeWrapReverse FlexWrapType = "wrap-reverse"
)

const (
	OverflowTypeVisible OverflowType = "visible"
	OverflowTypeHidden  OverflowType = "hidden"
	OverflowTypeScroll  OverflowType = "scroll"
	OverflowTypeAuto    OverflowType = "auto"
)

const (
	PositionTypeTopLeft     PositionType = "top-left"
	PositionTypeTopRight    PositionType = "top-right"
	PositionTypeBottomLeft  PositionType = "bottom-left"
	PositionTypeBottomRight PositionType = "bottom-right"
	PositionTypeTop         PositionType = "top"
	PositionTypeBottom      PositionType = "bottom"
	PositionTypeLeft        PositionType = "left"
	PositionTypeRight       PositionType = "right"
)

const (
	PositionStyleTypeFixed    PositionStyleType = "fixed"
	PositionStyleTypeAbsolute PositionStyleType = "absolute"
	PositionStyleTypeRelative PositionStyleType = "relative"
	PositionStyleTypeSticky   PositionStyleType = "sticky"
)

const (
	TextAlignTypeLeft    TextAlignType = "left"
	TextAlignTypeCenter  TextAlignType = "center"
	TextAlignTypeRight   TextAlignType = "right"
	TextAlignTypeJustify TextAlignType = "justify"
)

const (
	UserSelectTypeAll     UserSelectType = "all"
	UserSelectTypeText    UserSelectType = "text"
	UserSelectTypeNone    UserSelectType = "none"
	UserSelectTypeContain UserSelectType = "contain"
)
