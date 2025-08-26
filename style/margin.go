package style

import (
	"fmt"
	"strings"
)

type Margin struct {
	Top    int
	Right  int
	Bottom int
	Left   int
}

func (m *Margin) String() string {
	var (
		style = strings.Builder{}
	)

	if m.Top > 0 {
		style.WriteString(fmt.Sprintf("margin-top: %dpx;", m.Top))
	}
	if m.Right > 0 {
		style.WriteString(fmt.Sprintf("margin-right: %dpx;", m.Right))
	}
	if m.Bottom > 0 {
		style.WriteString(fmt.Sprintf("margin-bottom: %dpx;", m.Bottom))
	}
	if m.Left > 0 {
		style.WriteString(fmt.Sprintf("margin-left: %dpx;", m.Left))
	}

	return style.String()
}
