package options

type AxisAlignmentType string
type DisplayType string
type FlexDirectionType string
type AlignItemsType string
type AlignSelfType string
type AlignContentType string
type JustifyContentType string
type JustifyItemsType string
type JustifySelfType string
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
	DisplayTypeFlex DisplayType = "flex"
	DisplayTypeGrid DisplayType = "grid"
	DisplayTypeNone DisplayType = "none"
)

const (
	FlexDirectionTypeRow    FlexDirectionType = "row"
	FlexDirectionTypeColumn FlexDirectionType = "column"
)

const (
	AlignItemsTypeCenter   AlignItemsType = "center"
	AlignItemsTypeStart    AlignItemsType = "start"
	AlignItemsTypeEnd      AlignItemsType = "end"
	AlignItemsTypeStretch  AlignItemsType = "stretch"
	AlignItemsTypeBaseline AlignItemsType = "baseline"
)

const (
	AlignSelfTypeCenter   AlignSelfType = "center"
	AlignSelfTypeStart    AlignSelfType = "start"
	AlignSelfTypeEnd      AlignSelfType = "end"
	AlignSelfTypeStretch  AlignSelfType = "stretch"
	AlignSelfTypeBaseline AlignSelfType = "baseline"
)

const (
	AlignContentTypeCenter   AlignContentType = "center"
	AlignContentTypeStart    AlignContentType = "start"
	AlignContentTypeEnd      AlignContentType = "end"
	AlignContentTypeStretch  AlignContentType = "stretch"
	AlignContentTypeBaseline AlignContentType = "baseline"
)

const (
	JustifyContentTypeCenter   JustifyContentType = "center"
	JustifyContentTypeStart    JustifyContentType = "start"
	JustifyContentTypeEnd      JustifyContentType = "end"
	JustifyContentTypeStretch  JustifyContentType = "stretch"
	JustifyContentTypeBaseline JustifyContentType = "baseline"
)

const (
	JustifyItemsTypeCenter   JustifyItemsType = "center"
	JustifyItemsTypeStart    JustifyItemsType = "start"
	JustifyItemsTypeEnd      JustifyItemsType = "end"
	JustifyItemsTypeStretch  JustifyItemsType = "stretch"
	JustifyItemsTypeBaseline JustifyItemsType = "baseline"
)

const (
	JustifySelfTypeCenter   JustifySelfType = "center"
	JustifySelfTypeStart    JustifySelfType = "start"
	JustifySelfTypeEnd      JustifySelfType = "end"
	JustifySelfTypeStretch  JustifySelfType = "stretch"
	JustifySelfTypeBaseline JustifySelfType = "baseline"
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
