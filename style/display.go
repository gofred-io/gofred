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
	Display       DisplayType
	FlexDirection string
	Flex          int
	FlexWrap      string
}

func (d *Display) String() string {
	var (
		style = strings.Builder{}
	)

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

	return style.String()
}
