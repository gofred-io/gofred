package options

type AxisAlignmentType string
type BorderStyleType string
type DisplayType string
type FlexDirectionType string
type FlexWrapType string
type UserSelectType string

const (
	AxisAlignmentTypeCenter   AxisAlignmentType = "center"
	AxisAlignmentTypeStart    AxisAlignmentType = "start"
	AxisAlignmentTypeEnd      AxisAlignmentType = "end"
	AxisAlignmentTypeStretch  AxisAlignmentType = "stretch"
	AxisAlignmentTypeBaseline AxisAlignmentType = "baseline"
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
	UserSelectTypeAll     UserSelectType = "all"
	UserSelectTypeText    UserSelectType = "text"
	UserSelectTypeNone    UserSelectType = "none"
	UserSelectTypeContain UserSelectType = "contain"
)
