package style

import (
	"fmt"
	"strings"
)

type PaddingEx struct {
	Top    int
	Right  int
	Bottom int
	Left   int
}

func (p *PaddingEx) String() string {
	var (
		style = strings.Builder{}
	)

	if p.Top > 0 {
		style.WriteString(fmt.Sprintf("padding-top: %dpx;", p.Top))
	}
	if p.Right > 0 {
		style.WriteString(fmt.Sprintf("padding-right: %dpx;", p.Right))
	}
	if p.Bottom > 0 {
		style.WriteString(fmt.Sprintf("padding-bottom: %dpx;", p.Bottom))
	}
	if p.Left > 0 {
		style.WriteString(fmt.Sprintf("padding-left: %dpx;", p.Left))
	}

	return style.String()
}
