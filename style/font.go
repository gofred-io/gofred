package style

import (
	"fmt"
	"strings"
)

type Font struct {
	Color  string
	Family string
	Size   int
	Weight string
}

func (f *Font) String() string {
	var (
		style = strings.Builder{}
	)

	if f.Color != "" {
		style.WriteString(fmt.Sprintf("color: %s;", f.Color))
	}

	if f.Family != "" {
		style.WriteString(fmt.Sprintf("font-family: %s;", f.Family))
	}

	if f.Size > 0 {
		style.WriteString(fmt.Sprintf("font-size: %dpx;", f.Size))
	}

	if f.Weight != "" {
		style.WriteString(fmt.Sprintf("font-weight: %s;", f.Weight))
	}

	return style.String()
}
