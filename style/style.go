package style

import (
	"strings"
)

type Style struct {
	Background *Background
	Border     *Border
	Display    *Display
	Size       *Size
}

func (s *Style) String() string {
	var (
		style = strings.Builder{}
	)

	if s.Background != nil {
		style.WriteString(s.Background.String())
	}

	if s.Border != nil {
		style.WriteString(s.Border.String())
	}

	if s.Display != nil {
		style.WriteString(s.Display.String())
	}

	if s.Size != nil {
		style.WriteString(s.Size.String())
	}

	return style.String()
}
