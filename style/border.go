package style

import (
	"fmt"
	"strings"
)

type Border struct {
	Color  string
	Radius int
	Style  string
	Width  int
}

func (b *Border) String() string {
	var (
		style = strings.Builder{}
	)

	if b.Width > 0 {
		style.WriteString(fmt.Sprintf("border-width: %dpx;", b.Width))
	}

	if b.Color != "" {
		style.WriteString(fmt.Sprintf("border-color: %s;", b.Color))
	}

	if b.Radius > 0 {
		style.WriteString(fmt.Sprintf("border-radius: %dpx;", b.Radius))
	}

	if b.Style != "" {
		style.WriteString(fmt.Sprintf("border-style: %s;", b.Style))
	}

	return style.String()
}
