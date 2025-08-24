package style

import (
	"fmt"
	"strings"
)

type DisplayType string
type FlexDirectionType string
type AlignItemsType string
type AlignSelfType string
type AlignContent string
type JustifyContentType string
type JustifyItemsType string
type JustifySelfType string
type FlexWrapType string

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
	AlignContentTypeCenter   AlignContent = "center"
	AlignContentTypeStart    AlignContent = "start"
	AlignContentTypeEnd      AlignContent = "end"
	AlignContentTypeStretch  AlignContent = "stretch"
	AlignContentTypeBaseline AlignContent = "baseline"
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

type Display struct {
	AlignContent   AlignContent
	AlignItems     AlignItemsType
	AlignSelf      AlignSelfType
	Display        DisplayType
	FlexDirection  FlexDirectionType
	Flex           int
	FlexWrap       FlexWrapType
	JustifyContent JustifyContentType
	JustifyItems   JustifyItemsType
	JustifySelf    JustifySelfType
}

func (d *Display) String() string {
	var (
		style = strings.Builder{}
	)

	if d.AlignContent != "" {
		style.WriteString(fmt.Sprintf("align-content: %s;", d.AlignContent))
	}

	if d.AlignItems != "" {
		style.WriteString(fmt.Sprintf("align-items: %s;", d.AlignItems))
	}

	if d.AlignSelf != "" {
		style.WriteString(fmt.Sprintf("align-self: %s;", d.AlignSelf))
	}

	if d.Display != "" {
		style.WriteString(fmt.Sprintf("display: %s;", d.Display))
	}

	if d.FlexDirection != "" {
		style.WriteString(fmt.Sprintf("flex-direction: %s;", d.FlexDirection))
	}

	if d.Flex > 0 {
		style.WriteString(fmt.Sprintf("flex: %d;", d.Flex))
	}

	if d.FlexWrap != "" {
		style.WriteString(fmt.Sprintf("flex-wrap: %s;", d.FlexWrap))
	}

	if d.JustifyContent != "" {
		style.WriteString(fmt.Sprintf("justify-content: %s;", d.JustifyContent))
	}

	if d.JustifyItems != "" {
		style.WriteString(fmt.Sprintf("justify-items: %s;", d.JustifyItems))
	}

	if d.JustifySelf != "" {
		style.WriteString(fmt.Sprintf("justify-self: %s;", d.JustifySelf))
	}

	return style.String()
}
