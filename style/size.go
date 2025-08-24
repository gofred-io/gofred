package style

import (
	"fmt"
	"strings"
)

type Size struct {
	Width  int
	Height int
}

func (s *Size) String() string {
	var (
		style = strings.Builder{}
	)

	if s.Width > 0 {
		style.WriteString(fmt.Sprintf("width: %dpx;", s.Width))
	}

	if s.Height > 0 {
		style.WriteString(fmt.Sprintf("height: %dpx;", s.Height))
	}

	return style.String()
}
