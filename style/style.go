package style

import (
	"fmt"
	"strings"
)

type Style struct {
	Background *Background
	Border     *Border
	Display    *Display
	Fill       *string
	Font       *Font
	LineHeight *float64
	Margin     *Margin
	Padding    *Padding
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

	if s.Fill != nil {
		style.WriteString(fmt.Sprintf("fill: %s;", *s.Fill))
	}

	if s.Font != nil {
		style.WriteString(s.Font.String())
	}

	if s.LineHeight != nil {
		style.WriteString(fmt.Sprintf("line-height: %f;", *s.LineHeight))
	}

	if s.Padding != nil {
		style.WriteString(s.Padding.String())
	}

	if s.Size != nil {
		style.WriteString(s.Size.String())
	}

	return style.String()
}
