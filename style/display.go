package style

import (
	"fmt"
	"strings"
)

type DisplayType string

const (
	DisplayTypeFlex DisplayType = "flex"
	DisplayTypeGrid DisplayType = "grid"
	DisplayTypeNone DisplayType = "none"
)

type Display struct {
	AlignContent   string
	AlignItems     string
	AlignSelf      string
	Display        DisplayType
	FlexDirection  string
	Flex           int
	FlexWrap       string
	JustifyContent string
	JustifyItems   string
	JustifySelf    string
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
