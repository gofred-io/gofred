package style

import (
	"fmt"
	"strings"
)

type Size struct {
	Height    *int
	HeightP   *float32
	MaxHeight *int
	MaxWidth  *int
	Width     *int
	WidthP    *float32
}

func (s *Size) String() string {
	var (
		style = strings.Builder{}
	)

	if s.Width != nil {
		style.WriteString(fmt.Sprintf("width: %dpx;", *s.Width))
	} else if s.WidthP != nil {
		style.WriteString(fmt.Sprintf("width: %f%%;", *s.WidthP*100))
	}

	if s.Height != nil {
		style.WriteString(fmt.Sprintf("height: %dpx;", *s.Height))
	} else if s.HeightP != nil {
		style.WriteString(fmt.Sprintf("height: %f%%;", *s.HeightP*100))
	}

	if s.MaxWidth != nil {
		style.WriteString(fmt.Sprintf("max-width: %dpx;", *s.MaxWidth))
	}

	if s.MaxHeight != nil {
		style.WriteString(fmt.Sprintf("max-height: %dpx;", *s.MaxHeight))
	}

	return style.String()
}
